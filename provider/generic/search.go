package generic

import (
	"github.com/belphemur/mangal/source"
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

	err := s.mangasCollector.Visit(address)

	if err != nil {
		return nil, err
	}

	s.mangasCollector.Wait()

	_ = s.cache.mangas.Set(address, s.mangas)

	return s.mangas, nil
}
