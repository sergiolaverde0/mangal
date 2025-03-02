package none

import (
	"github.com/sergiolaverde0/mangal/constant"
	"github.com/sergiolaverde0/mangal/source"
)

type Converter struct {
}

func (c Converter) Format() (format constant.ConversionFormat) {
	return constant.ImageFormatNone
}

func (c Converter) ConvertChapter(chapter *source.Chapter, quality uint8, progress func(string)) (*source.Chapter, error) {
	return chapter, nil
}
