package webp

import (
	"bytes"
	"fmt"
	"github.com/metafates/mangal/key"
	"github.com/metafates/mangal/source"
	"github.com/oliamb/cutter"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
	"image"
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
	pagesChan := make(chan *source.Page, 5)
	var pages []*source.Page

	var wgPages sync.WaitGroup
	wgPages.Add(len(chapter.Pages))

	go func() {
		for page := range pagesChan {
			convertedPage, err := converter.convertPage(page)
			if err == nil {
				page = convertedPage
			}
			pages = append(pages, page)
			wgConvertedPages.Done()
		}
	}()

	for _, page := range chapter.Pages {
		go func(page *source.Page) {
			defer wgPages.Done()

			splitNeeded, img, err := converter.checkPageNeedsSplit(page)
			if err != nil {
				fmt.Println(err)
				return
			}
			if !splitNeeded {
				pagesChan <- page
				wgConvertedPages.Add(1)
				return
			}
			images, err := converter.cropImage(img)
			if err != nil {
				fmt.Println(err)
				return
			}

			for i, img := range images {
				buffer, err := convertImageToPngBytes(img)
				if err != nil {
					fmt.Println(err)
					return
				}

				page := &source.Page{Chapter: chapter, Index: page.Index, IsSplitted: true, SplitPartIndex: uint16(i), URL: page.URL, Extension: ".png", Contents: buffer, Size: uint64(buffer.Len())}
				pagesChan <- page
				wgConvertedPages.Add(1)
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

func (converter *Converter) checkPageNeedsSplit(page *source.Page) (bool, image.Image, error) {
	reader := io.Reader(bytes.NewBuffer(page.Contents.Bytes()))
	img, _, err := image.Decode(reader)
	if err != nil {
		return false, nil, err
	}
	bounds := img.Bounds()
	height := bounds.Dy()

	return height >= converter.maxHeight, img, nil
}

func (converter *Converter) convertPage(page *source.Page) (*source.Page, error) {
	converted, err := converter.convert(page.Contents.Bytes(), viper.GetUint(key.WebpQuality))
	if err != nil {
		return nil, err
	}
	page.Contents = converted
	page.Extension = ".webp"
	page.Size = uint64(converted.Len())
	return page, nil
}

// convert converts an image to the WebP format. It decodes the image from the input buffer,
// encodes it as a WebP file using the webp.Encode() function, and returns the resulting WebP
// file as a bytes.Buffer.
func (converter *Converter) convert(content []byte, quality uint) (*bytes.Buffer, error) {

	reader := io.Reader(bytes.NewBuffer(content))
	page, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = Encode(&buf, page, quality)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}

func convertImageToPngBytes(img image.Image) (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	err := png.Encode(buffer, img)
	if err != nil {
		return nil, fmt.Errorf("error writing image data to buffer: %v", err)
	}

	return buffer, nil
}
