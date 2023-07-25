package generic

import (
	"github.com/belphemur/mangal/source"
	"github.com/gocolly/colly/v2"
	"net/http"
)

// PagesOf given source.Chapter
func (s *Scraper) LoadPagesOf(chapter *source.Chapter) error {

	ctx := colly.NewContext()
	ctx.Put("chapter", chapter)
	err := s.pagesCollector.Request(http.MethodGet, chapter.URL, nil, ctx, nil)

	if err != nil {
		return err
	}

	s.pagesCollector.Wait()

	return nil
}
