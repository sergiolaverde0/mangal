package source

// Source is the interface that all sources must implement.
type Source interface {
	Name() string
	Search(query string) ([]*Manga, error)
	LoadChaptersOf(manga *Manga) error
	LoadPagesOf(chapter *Chapter) error
	ID() string
}
