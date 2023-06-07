package generic

import (
	"github.com/gocolly/colly/v2"
	"github.com/metafates/mangal/source"
	"net/http"
)

// ChaptersOf given source.Manga
func (s *Scraper) ChaptersOf(manga *source.Manga) ([]*source.Chapter, error) {
	if chapters := s.cache.chapters.Get(manga.URL); chapters.IsPresent() {
		c := chapters.MustGet()
		for _, chapter := range c {
			chapter.Manga = manga
		}

		return c, nil
	}

	if chapters, ok := s.chapters[manga.URL]; ok {
		return chapters, nil
	}

	ctx := colly.NewContext()
	ctx.Put("manga", manga)
	err := s.chaptersCollector.Request(http.MethodGet, manga.URL, nil, ctx, nil)

	if err != nil {
		return nil, err
	}

	s.chaptersCollector.Wait()

	if s.config.ReverseChapters {
		// reverse chapters
		chapters := s.chapters[manga.URL]
		reversed := make([]*source.Chapter, len(chapters))
		for i, chapter := range chapters {
			reversed[len(chapters)-i-1] = chapter
			chapter.Index = uint16(len(chapters) - i - 1)
			chapter.Index++
		}

		s.chapters[manga.URL] = reversed
	}
	_ = s.cache.chapters.Set(manga.URL, s.chapters[manga.URL])

	return s.chapters[manga.URL], nil
}
