package converter

import (
	"fmt"
	"github.com/belphemur/mangal/constant"
	"github.com/belphemur/mangal/converter/none"
	"github.com/belphemur/mangal/converter/webp"
	"github.com/belphemur/mangal/source"
	"github.com/samber/lo"
	"strings"
)

type Converter interface {
	// Format of the converter
	Format() (format constant.ConversionFormat)
	// ConvertChapter Conver the chapter in the given Format
	//  using the convertPage method. It runs the conversion concurrently for each page
	// with a maximum number of goroutines defined by maxGoroutines. The function returns the updated chapter
	// object after the conversion is complete.
	ConvertChapter(chapter *source.Chapter, quality uint8, progress func(string)) (*source.Chapter, error)
}

var converters = map[constant.ConversionFormat]Converter{
	constant.ImageFormatWebP: webp.New(),
	constant.ImageFormatNone: none.Converter{},
}

// Available returns a list of available converters.
func Available() []constant.ConversionFormat {
	return lo.Keys(converters)
}

// Get returns a packer by name.
// If the packer is not available, an error is returned.
func Get(name constant.ConversionFormat) (Converter, error) {
	if packer, ok := converters[name]; ok {
		return packer, nil
	}

	return nil, fmt.Errorf("unkown converter \"%s\", available options are %s", name, strings.Join(lo.Map(Available(), func(item constant.ConversionFormat, index int) string {
		return string(item)
	}), ", "))
}
