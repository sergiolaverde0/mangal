package generic

import (
	"github.com/belphemur/mangal/source"
	"github.com/gocolly/colly/v2"
	"net/http"
)

// ChaptersOf given source.Manga
func (s *Scraper) LoadChaptersOf(manga *source.Manga) error {
	if chapters := s.cache.chapters.Get(manga.URL); chapters.IsPresent() {
		c := chapters.MustGet()
		for _, chapter := range c {
			chapter.Manga = manga
		}

		return nil
	}

	ctx := colly.NewContext()
	ctx.Put("manga", manga)
	err := s.chaptersCollector.Request(http.MethodGet, manga.URL, nil, ctx, nil)

	if err != nil {
		return err
	}

	s.chaptersCollector.Wait()

	if s.config.ReverseChapters {
		// reverse chapters
		chapters := manga.Chapters
		reversed := make([]*source.Chapter, len(chapters))
		for i, chapter := range chapters {
			reversed[len(chapters)-i-1] = chapter
			chapter.Index = uint16(len(chapters) - i - 1)
			chapter.Index++
		}

		manga.Chapters = reversed
	}
	// Only cache if we have chapters
	if len(manga.Chapters) > 0 {
		_ = s.cache.chapters.Set(manga.URL, manga.Chapters)
	}

	return nil
}
