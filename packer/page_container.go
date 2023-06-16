package packer

import (
	"github.com/belphemur/mangal/source"
	"image"
)

// PageContainer is a struct that holds a manga page, its image, and the image format.
type PageContainer struct {
	// Page is a pointer to a manga page object.
	Page *source.Page
	// Image is the decoded image of the manga page.
	Image image.Image
	// Format is a string representing the format of the image (e.g., "png", "jpeg", "webp").
	Format string
}

func NewContainer(Page *source.Page, img image.Image, format string) *PageContainer {
	return &PageContainer{Page: Page, Image: img, Format: format}
}
