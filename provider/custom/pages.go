package custom

import (
	"github.com/sergiolaverde0/mangal/constant"
	"github.com/sergiolaverde0/mangal/source"
	lua "github.com/yuin/gopher-lua"
)

func (s *luaSource) LoadPagesOf(chapter *source.Chapter) error {
	_, err := s.call(constant.ChapterPagesFn, lua.LTTable, lua.LString(chapter.URL))

	if err != nil {
		return err
	}

	table := s.state.CheckTable(-1)
	pages := make([]*source.Page, 0)

	table.ForEach(func(k lua.LValue, v lua.LValue) {
		if k.Type() != lua.LTNumber {
			s.state.RaiseError(constant.ChapterPagesFn + " was expected to return a table with numbers as keys, got " + k.Type().String() + " as a key")
		}

		if v.Type() != lua.LTTable {
			s.state.RaiseError(constant.ChapterPagesFn + " was expected to return a table with tables as values, got " + v.Type().String() + " as a value")
		}

		page, err := pageFromTable(v.(*lua.LTable), chapter)

		if err != nil {
			s.state.RaiseError(err.Error())
		}

		pages = append(pages, page)
	})

	chapter.Pages = pages

	return nil
}
