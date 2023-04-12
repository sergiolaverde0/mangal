package webp

import (
	"bytes"
	"github.com/chai2010/webp"
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

// CheckAndConvert checks whether WebP conversion is enabled in the configuration, and if so,
// invokes the convert method to convert the image to WebP format. If WebP conversion is disabled,
// it returns the original image content without modification.
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
	err = webp.Encode(&buf, page, &webp.Options{Quality: float32(quality)})
	if err != nil {
		return nil, err
	}

	return &buf, nil
}
