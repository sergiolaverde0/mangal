package generic

import (
	"github.com/belphemur/mangal/source"
	"github.com/gocolly/colly/v2"
)

// Search for mangas by given title
func (s *Scraper) Search(query string) ([]*source.Manga, error) {
	address := s.config.GenerateSearchURL(s.config.BaseURL, query)

	if mangas := s.cache.mangas.Get(address); mangas.IsPresent() {
		m := mangas.MustGet()
		for _, manga := range m {
			manga.Source = s
		}

		return m, nil
	}

	ctx := colly.NewContext()
	collector := &MangaResult{Mangas: make([]*source.Manga, 0)}
	ctx.Put("collector", collector)

	err := s.mangasCollector.Request("GET", address, nil, ctx, nil)

	if err != nil {
		return nil, err
	}

	s.mangasCollector.Wait()

	_ = s.cache.mangas.Set(address, collector.Mangas)

	return collector.Mangas, nil
}
