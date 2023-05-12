package none

import (
	"github.com/metafates/mangal/constant"
	"github.com/metafates/mangal/source"
)

type Converter struct {
}

func (c Converter) Format() (format constant.ConversionFormat) {
	return constant.ImageFormatNone
}

func (c Converter) ConvertChapter(chapter *source.Chapter) (*source.Chapter, error) {
	return chapter, nil
}
