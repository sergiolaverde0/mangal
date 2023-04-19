package webp

import (
	"bytes"
	"github.com/metafates/mangal/key"
	"github.com/metafates/mangal/source"
	"github.com/spf13/viper"
	"image"
	"io"
	"sync"
)

type Converter struct{}

func New() *Converter {
	return &Converter{}
}

// CheckAndConvertChapter checks if WebP conversion is enabled and converts each page in the chapter
// to WebP format using the convertPage method. It runs the conversion concurrently for each page
// with a maximum number of goroutines defined by maxGoroutines. The function returns the updated chapter
// object after the conversion is complete.
func (converter *Converter) CheckAndConvertChapter(chapter *source.Chapter) (*source.Chapter, error) {
	if !viper.GetBool(key.WebpConversion) {
		return chapter, nil
	}
	var wg sync.WaitGroup
	maxGoroutines := 5
	guard := make(chan struct{}, maxGoroutines)
	wg.Add(len(chapter.Pages))

	for i, page := range chapter.Pages {
		guard <- struct{}{} // would block if guard channel is already filled
		go func(index int, page *source.Page) {
			defer wg.Done()
			convertedPage, err := converter.convertPage(page)
			if err == nil {
				page = convertedPage
			}
			chapter.Pages[index] = page

			<-guard
		}(i, page)
	}
	wg.Wait()
	return chapter, nil
}

// CheckAndConvert checks if the WebP conversion is enabled in the configuration and
// converts the page contents to WebP format if it is. It uses the Converter's convert
// method to perform the conversion with the specified WebP quality. It then updates the
// page's contents and extension to reflect the conversion.
//
// If the conversion is disabled, the function returns the original page without any
// modifications.
//
// Returns the modified page with the WebP contents and extension, or an error if the
// conversion fails.
func (converter *Converter) CheckAndConvert(page *source.Page) (*source.Page, error) {
	if !viper.GetBool(key.WebpConversion) {
		return page, nil
	}

	return converter.convertPage(page)
}

func (converter *Converter) convertPage(page *source.Page) (*source.Page, error) {
	converted, err := converter.convert(page.Contents, viper.GetUint(key.WebpQuality))
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
func (converter *Converter) convert(content *bytes.Buffer, quality uint) (*bytes.Buffer, error) {

	reader := io.Reader(content)
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
