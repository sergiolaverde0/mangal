package converter

import (
	"github.com/metafates/mangal/constant"
	"github.com/metafates/mangal/source"
)

type Converter interface {
	// Format of the converter
	Format() (format constant.ConversionFormat)
	// CheckAndConvertChapter checks if WebP conversion is enabled and converts each page in the chapter
	// to WebP format using the convertPage method. It runs the conversion concurrently for each page
	// with a maximum number of goroutines defined by maxGoroutines. The function returns the updated chapter
	// object after the conversion is complete.
	CheckAndConvertChapter(chapter *source.Chapter) (*source.Chapter, error)
}
