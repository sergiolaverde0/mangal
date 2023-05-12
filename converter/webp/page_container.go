package webp

import (
	"github.com/metafates/mangal/source"
	"image"
)

type PageContainer struct {
	Page  *source.Page
	Image *image.Image
}

func NewContainer(page *source.Page) *PageContainer {
	decoded, _, err := image.Decode(page.Contents)
	if err != nil {
		return nil
	}
	return &PageContainer{page, &decoded}
}

func NewContainerWithImage(Page *source.Page, img *image.Image) *PageContainer {
	return &PageContainer{Page: Page, Image: img}
}
