package webp

import (
	"bytes"
	"github.com/metafates/mangal/key"
	"github.com/metafates/mangal/source"
	"github.com/spf13/viper"
	"image"
	"io"
)

type Converter struct{}

func New() *Converter {
	return &Converter{}
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
