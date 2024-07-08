package tui

import (
	"fmt"
	"github.com/belphemur/mangal/anilist"
	"github.com/belphemur/mangal/history"
	"github.com/belphemur/mangal/icon"
	"github.com/belphemur/mangal/installer"
	"github.com/belphemur/mangal/provider"
	"github.com/belphemur/mangal/source"
	"github.com/belphemur/mangal/style"
	"strings"
)

type listItem struct {
	internal interface{}
	marked   bool
}

func (t *listItem) toggleMark() {
	t.marked = !t.marked
}

func (t *listItem) getMark() string {
	switch t.internal.(type) {
	case *source.Chapter:
		return style.Bold(icon.Get(icon.Mark))
	case *anilist.Manga:
		return icon.Get(icon.Link)
	case *provider.Provider:
		return icon.Get(icon.Search)
	default:
		return ""
	}
}

func (t *listItem) Title() (title string) {
	switch e := t.internal.(type) {
	case *source.Chapter:
		var sb = strings.Builder{}

		sb.WriteString(t.FilterValue())
		if e.Volume != "" {
			sb.WriteString(" ")
			sb.WriteString(style.Faint(e.Volume))
		}

		if e.ChapterDate != nil {
			sb.WriteString(style.Faint(fmt.Sprintf(" [%s]", e.ChapterDate.Format("2006-01-02"))))
		}

		if e.IsDownloaded() {
			sb.WriteString(" ")
			sb.WriteString(icon.Get(icon.Downloaded))
		}

		title = sb.String()
	default:
		title = t.FilterValue()
	}

	if title != "" && t.marked {
		//title = fmt.Sprintf("%s %s", title, icon.Get(icon.Mark))
		title = fmt.Sprintf("%s %s", title, t.getMark())
	}

	return
}

func (t *listItem) Description() (description string) {
	switch e := t.internal.(type) {
	case *source.Chapter:
		var sb = strings.Builder{}
		if len(e.Scanlations) != 0 {
			sb.WriteString(style.Italic(strings.Join(e.Scanlations, ", ")))
			sb.WriteString(": ")
		}
		sb.WriteString(e.URL)

		description = sb.String()
	case *source.Manga:
		description = e.URL
	case *installer.Scraper:
		description = e.GithubURL()
	case *history.SavedChapter:
		description = fmt.Sprintf("%s : %d / %d", e.Name, e.Index, e.MangaChaptersTotal)
	case *provider.Provider:
		sb := strings.Builder{}
		if e.IsCustom {
			sb.WriteString("Custom")
		} else {
			sb.WriteString("Builtin")
		}

		if e.UsesHeadless {
			sb.WriteString(", uses headless chrome")
		}

		description = sb.String()
	case *anilist.Manga:
		description = e.SiteURL
	}

	return
}

func (t *listItem) FilterValue() string {
	switch e := t.internal.(type) {
	case *source.Chapter:
		return e.Name
	case *source.Manga:
		return e.Name
	case *history.SavedChapter:
		return e.MangaName
	case *anilist.Manga:
		return e.Name()
	case *provider.Provider:
		return e.Name
	case *installer.Scraper:
		return e.Name
	default:
		return ""
	}
}
