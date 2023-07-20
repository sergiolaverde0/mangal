package asurascans

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/belphemur/mangal/provider/generic"
	"strings"
	"time"
)

const dateLayout = "January 2, 2006"

var Config = &generic.Configuration{
	Name:                 "AsuraScans",
	Delay:                50 * time.Millisecond,
	Parallelism:          15,
	ReverseChapters:      true,
	NeedsHeadlessBrowser: true,
	BaseURL:              "https://asura.gg",
	GenerateSearchURL: func(baseUrl string, query string) string {
		query = strings.ReplaceAll(query, "’s", "")
		query = strings.ReplaceAll(query, "'s", "")
		query = strings.ReplaceAll(query, "’ll", "")
		query = strings.ReplaceAll(query, "'ll", "")
		return baseUrl + "/?s=" + query
	},
	MangaExtractor: &generic.MangaExtractor{
		Selector: ".bsx > a",
		Name: func(selection *goquery.Selection) string {
			return strings.TrimSpace(selection.AttrOr("title", ""))
		},
		URL: func(selection *goquery.Selection) string {
			return selection.AttrOr("href", "")
		},
		Cover: func(selection *goquery.Selection) string {
			return selection.Find("img").AttrOr("src", "")
		},
	},
	ChapterExtractor: &generic.ChapterExtractor{
		Selector: "#chapterlist > ul li",
		Name: func(selection *goquery.Selection) string {
			name := selection.Find(".chapternum").Text()
			return name
		},
		URL: func(selection *goquery.Selection) string {
			return selection.Find("a").AttrOr("href", "")
		},
		Volume: func(selection *goquery.Selection) string {
			name := selection.Find(".chapternum").Text()
			if strings.HasPrefix(name, "Vol.") {
				splitted := strings.Split(name, " ")
				return splitted[0]
			}
			return ""
		},
		Date: func(selection *goquery.Selection) *time.Time {
			date := selection.Find(".chapterdate").Text()
			t, err := time.Parse(dateLayout, date)
			if err != nil {
				return nil
			}
			return &t
		},
	},
	PageExtractor: &generic.PageExtractor{
		Selector: "#readerarea img",
		URL: func(selection *goquery.Selection) string {
			return selection.AttrOr("src", "")
		},
	},
}
