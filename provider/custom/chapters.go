package custom

import (
	"github.com/sergiolaverde0/mangal/constant"
	"github.com/sergiolaverde0/mangal/source"
	lua "github.com/yuin/gopher-lua"
	"golang.org/x/exp/slices"
	"strconv"
)

func (s *luaSource) LoadChaptersOf(manga *source.Manga) error {
	if chapters := s.cache.chapters.Get(manga.URL); chapters.IsPresent() {
		c := chapters.MustGet()
		for _, chapter := range c {
			chapter.Manga = manga
		}

		return nil
	}

	_, err := s.call(constant.MangaChaptersFn, lua.LTTable, lua.LString(manga.URL))

	if err != nil {
		return err
	}

	table := s.state.CheckTable(-1)
	chapters := make([]*source.Chapter, 0)

	table.ForEach(func(k lua.LValue, v lua.LValue) {
		if k.Type() != lua.LTNumber {
			s.state.RaiseError(constant.MangaChaptersFn + " was expected to return a table with numbers as keys, got " + k.Type().String() + " as a key")
		}

		if v.Type() != lua.LTTable {
			s.state.RaiseError(constant.MangaChaptersFn + " was expected to return a table with tables as values, got " + v.Type().String() + " as a value")
		}

		index, err := strconv.ParseUint(k.String(), 10, 16)
		if err != nil {
			s.state.RaiseError(constant.MangaChaptersFn + " was expected to return a table with unsigned integers as keys. " + err.Error())
		}

		chapter, err := chapterFromTable(v.(*lua.LTable), manga, uint16(index))

		if err != nil {
			s.state.RaiseError(err.Error())
		}

		chapters = append(chapters, chapter)
	})
	slices.SortFunc(chapters, func(a, b *source.Chapter) int {
		return int(b.Number - a.Number)
	})

	for i, chapter := range chapters {
		chapter.Index = uint16(i + 1)
	}

	_ = s.cache.chapters.Set(manga.URL, chapters)
	manga.Chapters = chapters
	return nil
}
