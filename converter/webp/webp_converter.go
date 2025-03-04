package webp

import (
	"bytes"
	"fmt"
	"github.com/sergiolaverde0/mangal/constant"
	"github.com/sergiolaverde0/mangal/log"
	"github.com/sergiolaverde0/mangal/packer"
	"github.com/sergiolaverde0/mangal/source"
	"github.com/oliamb/cutter"
	"golang.org/x/exp/slices"
	_ "golang.org/x/image/webp"
	"image"
	_ "image/jpeg"
	"image/png"
	"io"
	"runtime"
	"sync"
	"sync/atomic"
)

type Converter struct {
	maxHeight  int
	cropHeight int
}

func (converter *Converter) Format() (format constant.ConversionFormat) {
	return constant.ImageFormatWebP
}

func New() *Converter {
	return &Converter{
		//maxHeight: 16383 / 2,
		maxHeight:  4000,
		cropHeight: 2000,
	}
}

func (converter *Converter) ConvertChapter(chapter *source.Chapter, quality uint8, progress func(string)) (*source.Chapter, error) {
	err := PrepareEncoder()
	if err != nil {
		return nil, err
	}

	var wgConvertedPages sync.WaitGroup
	maxGoroutines := runtime.NumCPU()

	pagesChan := make(chan *packer.PageContainer, maxGoroutines)

	var wgPages sync.WaitGroup
	wgPages.Add(len(chapter.Pages))

	guard := make(chan struct{}, maxGoroutines)
	pagesMutex := sync.Mutex{}
	var pages []*source.Page
	var totalPages = uint32(len(chapter.Pages))

	go func() {
		for page := range pagesChan {
			guard <- struct{}{} // would block if guard channel is already filled
			go func(pageToConvert *packer.PageContainer) {
				defer wgConvertedPages.Done()
				convertedPage, err := converter.convertPage(pageToConvert, quality)
				if err != nil {
					buffer := new(bytes.Buffer)
					err := png.Encode(buffer, convertedPage.Image)
					if err != nil {
						<-guard
						return
					}
					convertedPage.Page.Contents = buffer
					convertedPage.Page.Extension = ".png"
					convertedPage.Page.Size = uint64(buffer.Len())
				}
				pagesMutex.Lock()
				pages = append(pages, convertedPage.Page)
				progress(fmt.Sprintf("Converted %d/%d pages to %s format", len(pages), totalPages, converter.Format()))
				pagesMutex.Unlock()
				<-guard
			}(page)

		}
	}()

	for _, page := range chapter.Pages {
		go func(page *source.Page) {
			defer wgPages.Done()

			splitNeeded, img, format, err := converter.checkPageNeedsSplit(page)
			if err != nil {
				log.Error(fmt.Sprintf("error checking if page %d of manga %s needs split: %v", page.Index, page.Chapter.Manga.Name, err))
				return
			}

			if !splitNeeded {
				wgConvertedPages.Add(1)
				pagesChan <- packer.NewContainer(page, img, format)
				return
			}
			images, err := converter.cropImage(img)
			if err != nil {
				log.Error(fmt.Sprintf("error converting page %d of manga %s to webp: %v", page.Index, page.Chapter.Manga.Name, err))
				return
			}

			atomic.AddUint32(&totalPages, uint32(len(images)-1))
			for i, img := range images {
				page := &source.Page{Chapter: chapter, Index: page.Index, IsSplitted: true, SplitPartIndex: uint16(i), URL: page.URL}
				wgConvertedPages.Add(1)
				pagesChan <- packer.NewContainer(page, img, "N/A")
			}
		}(page)

	}

	wgPages.Wait()
	wgConvertedPages.Wait()
	close(pagesChan)

	slices.SortFunc(pages, func(a, b *source.Page) int {
		if a.Index == b.Index {
			return int(b.SplitPartIndex - a.SplitPartIndex)
		}
		return int(b.Index - a.Index)
	})
	chapter.UpdatePages(pages)

	runtime.GC()

	return chapter, nil
}

func (converter *Converter) cropImage(img image.Image) ([]image.Image, error) {
	bounds := img.Bounds()
	height := bounds.Dy()

	numParts := height / converter.cropHeight
	if height%converter.cropHeight != 0 {
		numParts++
	}

	parts := make([]image.Image, numParts)

	for i := 0; i < numParts; i++ {
		partHeight := converter.cropHeight
		if i == numParts-1 {
			partHeight = height - i*converter.cropHeight
		}

		part, err := cutter.Crop(img, cutter.Config{
			Width:  bounds.Dx(),
			Height: partHeight,
			Anchor: image.Point{0, i * converter.cropHeight},
			Mode:   cutter.TopLeft,
		})
		if err != nil {
			return nil, fmt.Errorf("error cropping part %d: %v", i+1, err)
		}

		parts[i] = part
	}

	return parts, nil
}

func (converter *Converter) checkPageNeedsSplit(page *source.Page) (bool, image.Image, string, error) {
	reader := io.Reader(bytes.NewBuffer(page.Contents.Bytes()))
	img, format, err := image.Decode(reader)
	if err != nil {
		return false, nil, format, err
	}

	bounds := img.Bounds()
	height := bounds.Dy()

	return height >= converter.maxHeight, img, format, nil
}

func (converter *Converter) convertPage(container *packer.PageContainer, quality uint8) (*packer.PageContainer, error) {
	if container.Format == "webp" {
		return container, nil
	}
	converted, err := converter.convert(container.Image, uint(quality))
	if err != nil {
		return nil, err
	}
	container.Page.Contents = converted
	container.Page.Extension = ".webp"
	container.Page.Size = uint64(converted.Len())
	return container, nil
}

// convert converts an image to the ImageFormatWebP format. It decodes the image from the input buffer,
// encodes it as a ImageFormatWebP file using the webp.Encode() function, and returns the resulting ImageFormatWebP
// file as a bytes.Buffer.
func (converter *Converter) convert(image image.Image, quality uint) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	err := Encode(&buf, image, quality)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}
