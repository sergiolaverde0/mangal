package webp

import (
	"bytes"
	"fmt"
	"github.com/metafates/mangal/key"
	"github.com/metafates/mangal/source"
	"github.com/oliamb/cutter"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
	_ "golang.org/x/image/webp"
	"image"
	_ "image/jpeg"
	"image/png"
	"io"
	"sync"
)

type Converter struct {
	maxHeight int
}

func New() *Converter {
	return &Converter{
		maxHeight: 16383 / 2,
	}
}

// CheckAndConvertChapter checks if WebP conversion is enabled and converts each page in the chapter
// to WebP format using the convertPage method. It runs the conversion concurrently for each page
// with a maximum number of goroutines defined by maxGoroutines. The function returns the updated chapter
// object after the conversion is complete.
func (converter *Converter) CheckAndConvertChapter(chapter *source.Chapter) (*source.Chapter, error) {
	if !viper.GetBool(key.WebpConversion) {
		return chapter, nil
	}

	var wgConvertedPages sync.WaitGroup
	maxGoroutines := 6

	pagesChan := make(chan *PageContainer, maxGoroutines*2)

	var wgPages sync.WaitGroup
	wgPages.Add(len(chapter.Pages))

	guard := make(chan struct{}, maxGoroutines)
	pagesMutex := sync.Mutex{}
	var pages []*source.Page

	go func() {
		for page := range pagesChan {
			guard <- struct{}{} // would block if guard channel is already filled
			go func(pageToConvert *PageContainer) {
				defer wgConvertedPages.Done()
				convertedPage, err := converter.convertPage(pageToConvert)
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
				pagesChan <- NewContainer(page, img, format)
				return
			}
			images, err := converter.cropImage(img)
			if err != nil {
				fmt.Println(err)
				return
			}

			for i, img := range images {
				page := &source.Page{Chapter: chapter, Index: page.Index, IsSplitted: true, SplitPartIndex: uint16(i), URL: page.URL, Extension: page.Extension, Contents: page.Contents, Size: page.Size}
				wgConvertedPages.Add(1)
				pagesChan <- NewContainer(page, img, "N/A")
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
	chapter.Pages = pages

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

func (converter *Converter) convertPage(container *PageContainer) (*PageContainer, error) {
	if container.Format == "webp" {
		return container, nil
	}
	converted, err := converter.convert(container.Image, viper.GetUint(key.WebpQuality))
	if err != nil {
		return nil, err
	}
	container.Page.Contents = converted
	container.Page.Extension = ".webp"
	container.Page.Size = uint64(converted.Len())
	return container, nil
}

// convert converts an image to the WebP format. It decodes the image from the input buffer,
// encodes it as a WebP file using the webp.Encode() function, and returns the resulting WebP
// file as a bytes.Buffer.
func (converter *Converter) convert(image image.Image, quality uint) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	err := Encode(&buf, image, quality)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}
