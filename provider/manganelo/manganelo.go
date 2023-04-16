package manganelo

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/metafates/mangal/provider/generic"
	"github.com/metafates/mangal/util"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var hoursAgoRegex = regexp.MustCompile(`(?m)(?P<hours>\d{1,2}) hour ago`)
var Config = &generic.Configuration{
	Name:            "Manganelo",
	Delay:           50 * time.Millisecond,
	Parallelism:     50,
	ReverseChapters: true,
	BaseURL:         "https://ww5.manganelo.tv/",
	GenerateSearchURL: func(query string) string {
		query = strings.TrimSpace(query)
		query = strings.ToLower(query)
		query = url.QueryEscape(query)
		template := "https://ww5.manganelo.tv/search/%s"
		return fmt.Sprintf(template, query)
	},
	MangaExtractor: &generic.MangaExtractor{
		Selector: ".search-story-item",
		Name: func(selection *goquery.Selection) string {
			return selection.Find("a.item-title").Text()
		},
		URL: func(selection *goquery.Selection) string {
			return selection.Find("a.item-title").AttrOr("href", "")
		},
		Cover: func(selection *goquery.Selection) string {
			return selection.Find(".item-img img").AttrOr("src", "")
		},
	},
	ChapterExtractor: &generic.ChapterExtractor{
		Selector: "li.a-h",
		Name: func(selection *goquery.Selection) string {
			name := selection.Find(".chapter-name").Text()
			if strings.HasPrefix(name, "Vol.") {
				splitted := strings.Split(name, " ")
				name = strings.Join(splitted[1:], " ")
			}
			return name
		},
		URL: func(selection *goquery.Selection) string {
			return selection.Find(".chapter-name").AttrOr("href", "")
		},
		Volume: func(selection *goquery.Selection) string {
			name := selection.Find(".chapter-name").Text()
			if strings.HasPrefix(name, "Vol.") {
				splitted := strings.Split(name, " ")
				return splitted[0]
			}
			return ""
		},
		Date: func(selection *goquery.Selection) *time.Time {
			layout := "Jan 02,06"
			publishedDate := strings.TrimSpace(selection.Find(".chapter-time").Text())
			date, err := time.Parse(layout, publishedDate)
			if err != nil {
				dateMatch := util.ReGroups(hoursAgoRegex, publishedDate)
				if len(dateMatch) == 0 {
					return nil
				}
				hours, err := strconv.ParseInt(dateMatch["hours"], 10, 8)
				if err != nil {
					return nil
				}
				date = time.Now().Add(time.Duration(-hours) * time.Hour)
			}
			return &date
		},
	},
	PageExtractor: &generic.PageExtractor{
		Selector: ".container-chapter-reader img",
		Name:     nil,
		URL: func(selection *goquery.Selection) string {
			return selection.AttrOr("data-src", "")
		},
	},
}
