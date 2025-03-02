package asurascans

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/araddon/dateparse"
	"github.com/sergiolaverde0/mangal/provider/generic"
	"strings"
	"time"
)

var Config = &generic.Configuration{
	Name:                 "AsuraScans",
	Delay:                50 * time.Millisecond,
	Parallelism:          15,
	ReverseChapters:      true,
	NeedsHeadlessBrowser: true,
	BaseURL:              "https://asuracomic.net/",
	GenerateSearchURL: func(baseUrl string, query string) string {
		query = strings.ReplaceAll(query, "’s", "")
		query = strings.ReplaceAll(query, "'s", "")
		query = strings.ReplaceAll(query, "’ll", "")
		query = strings.ReplaceAll(query, "'ll", "")
		return baseUrl + "series?page=1&name=" + query
	},
	MangaExtractor: &generic.MangaExtractor{
		Selector: ".grid > a",
		Name: func(selection *goquery.Selection) string {
			return strings.TrimSpace(selection.Find("div:nth-of-type(2) > span:nth-of-type(1)").Text())
		},
		URL: func(selection *goquery.Selection) string {
			return selection.AttrOr("href", "")
		},
		Cover: func(selection *goquery.Selection) string {
			return selection.Find("img").AttrOr("src", "")
		},
	},
	ChapterExtractor: &generic.ChapterExtractor{
		Selector: "div.pl-4.py-2.border.rounded-md.group.w-full",
		Name: func(selection *goquery.Selection) string {
			name := selection.Find("h3:nth-of-type(1)").Text()
			return name
		},
		URL: func(selection *goquery.Selection) string {
			return selection.Find("a").AttrOr("href", "")
		},
		Volume: func(selection *goquery.Selection) string {
			name := selection.Find("h3:nth-of-type(1)").Text()
			if strings.HasPrefix(name, "Vol.") {
				splitted := strings.Split(name, " ")
				return splitted[0]
			}
			return ""
		},
		Date: func(selection *goquery.Selection) *time.Time {
			date := selection.Find("h3:nth-of-type(2)").Text()
			t, err := dateparse.ParseAny(date)
			if err != nil {
				return nil
			}
			return &t
		},
	},
	PageExtractor: &generic.PageExtractor{
		Selector: "div.w-full > img.mx-auto",
		URL: func(selection *goquery.Selection) string {
			return selection.AttrOr("src", "")
		},
	},
}
