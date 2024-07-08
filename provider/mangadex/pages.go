package mangadex

import (
	"bytes"
	"errors"
	"github.com/belphemur/mangal/source"
	"path/filepath"
)

func (m *Mangadex) LoadPagesOf(chapter *source.Chapter) error {
	downloader, err := m.client.AtHome.NewMDHomeClient(chapter.ID, "data", false)
	if err != nil {
		return err
	}

	if len(downloader.Pages) == 0 {
		return errors.New("there were no pages for this chapter")
	}

	var pages = make([]*source.Page, len(downloader.Pages))

	for i, name := range downloader.Pages {
		image, err := downloader.GetChapterPage(name)
		if err != nil {
			return err
		}

		if len(image) == 0 {
			return errors.New("image is empty")
		}

		page := source.Page{
			Index:     uint16(i),
			Chapter:   chapter,
			Extension: filepath.Ext(name),
			Contents:  bytes.NewBuffer(image),
			Size:      uint64(len(image)),
		}

		pages[i] = &page
	}

	chapter.Pages = pages
	return nil
}
