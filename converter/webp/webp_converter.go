package webp

import (
	"bytes"
	"fmt"
	"github.com/metafates/mangal/constant"
	"github.com/metafates/mangal/packer"
	"github.com/metafates/mangal/source"
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
	maxHeight int
}

func (converter *Converter) Format() (format constant.ConversionFormat) {
	return constant.ImageFormatWebP
}

func New() *Converter {
	return &Converter{
		maxHeight: 16383 / 2,
	}
}

func (converter *Converter) ConvertChapter(chapter *source.Chapter, quality uint8, progress func(string)) (*source.Chapter, error) {
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
					err := png.Encode(buffer, pageToConvert.Image)
					if err != nil {
						<-guard
						return
					}
					pageToConvert.Page.Contents = buffer
					pageToConvert.Page.Extension = ".png"
					pageToConvert.Page.Size = uint64(buffer.Len())
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
				fmt.Println(err)
				return
			}

			if !splitNeeded {
				wgConvertedPages.Add(1)
				pagesChan <- packer.NewContainer(page, img, format)
				return
			}
			images, err := converter.cropImage(img)
			if err != nil {
				fmt.Println(err)
				return
			}

			atomic.AddUint32(&totalPages, uint32(len(images)-1))
			for i, img := range images {
				page := &source.Page{Chapter: chapter, Index: page.Index, IsSplitted: true, SplitPartIndex: uint16(i), URL: page.URL, Extension: page.Extension, Contents: page.Contents, Size: page.Size}
				wgConvertedPages.Add(1)
				pagesChan <- packer.NewContainer(page, img, "N/A")
			}
		}(page)

	}

	wgPages.Wait()
	wgConvertedPages.Wait()
	close(pagesChan)

	slices.SortFunc(pages, func(a, b *source.Page) bool {
		if a.Index == b.Index {
			return a.SplitPartIndex < b.SplitPartIndex
		}
		return a.Index < b.Index
	})
	chapter.UpdatePages(pages)

	return chapter, nil
}

func (converter *Converter) cropImage(img image.Image) ([]image.Image, error) {
	bounds := img.Bounds()
	height := bounds.Dy()

	numParts := height / converter.maxHeight
	if height%converter.maxHeight != 0 {
		numParts++
	}

	parts := make([]image.Image, numParts)

	for i := 0; i < numParts; i++ {
		partHeight := converter.maxHeight
		if i == numParts-1 {
			partHeight = height - i*converter.maxHeight
		}

		part, err := cutter.Crop(img, cutter.Config{
			Width:  bounds.Dx(),
			Height: partHeight,
			Anchor: image.Point{0, i * converter.maxHeight},
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
	if format == "webp" {
		return false, img, format, nil
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
