package generic

import (
	"github.com/belphemur/mangal/source"
	"github.com/gocolly/colly/v2"
	"net/http"
)

// PagesOf given source.Chapter
func (s *Scraper) PagesOf(chapter *source.Chapter) ([]*source.Page, error) {

	ctx := colly.NewContext()
	ctx.Put("chapter", chapter)
	err := s.pagesCollector.Request(http.MethodGet, chapter.URL, nil, ctx, nil)

	if err != nil {
		return nil, err
	}

	s.pagesCollector.Wait()

	return s.pages, nil
}
