package mangadex

import (
	"fmt"
	"github.com/darylhjd/mangodex"
	"github.com/metafates/mangal/key"
	"github.com/metafates/mangal/source"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
	"net/url"
	"strconv"
	"time"
)

func removeDuplicate(chapters []*source.Chapter) ([]*source.Chapter, error) {
	if !viper.GetBool(key.MangadexAvoidDuplicateChapters) {
		return chapters, nil
	}

	uniqueChapters := lo.UniqBy(chapters, func(item *source.Chapter) float32 {
		return item.Number
	})

	return lo.Map(uniqueChapters, func(item *source.Chapter, index int) *source.Chapter {
		item.Index = uint16(index + 1)
		return item
	}), nil
}
func (m *Mangadex) ChaptersOf(manga *source.Manga) ([]*source.Chapter, error) {
	if cached, ok := m.cache.chapters.Get(manga.URL).Get(); ok {
		for _, chapter := range cached {
			chapter.Manga = manga
		}

		return cached, nil
	}
	avoidDuplicatedChapters := viper.GetBool(key.MangadexAvoidDuplicateChapters)

	params := url.Values{}
	params.Set("limit", strconv.Itoa(500))
	ratings := []string{mangodex.Safe, mangodex.Suggestive}
	for _, rating := range ratings {
		params.Add("contentRating[]", rating)
	}

	if viper.GetBool(key.MangadexNSFW) {
		params.Add("contentRating[]", mangodex.Porn)
		params.Add("contentRating[]", mangodex.Erotica)
	}

	// scanlation group for the chapter
	params.Add("includes[]", mangodex.ScanlationGroupRel)
	params.Set("order[chapter]", "asc")

	var (
		chapters   []*source.Chapter
		currOffset = 0
	)

	language := viper.GetString(key.MangadexLanguage)

	var chapterIndex uint16 = 1
	chapterSet := make(map[float32]bool)

	for {
		params.Set("offset", strconv.Itoa(currOffset))
		list, err := m.client.Chapter.GetMangaChapters(manga.ID, params)
		if err != nil {
			return nil, err
		}

		for _, chapter := range list.Data {
			// Skip external chapters. Their pages cannot be downloaded.
			if chapter.Attributes.ExternalURL != nil && !viper.GetBool(key.MangadexShowUnavailableChapters) {
				continue
			}

			// skip chapters that are not in the current language
			if language != "any" && chapter.Attributes.TranslatedLanguage != language {
				currOffset += 500
				continue
			}

			number, err := strconv.ParseFloat(chapter.GetChapterNum(), 32)
			if err != nil {
				number = 0
			}
			chapterNumber := float32(number)

			if avoidDuplicatedChapters && chapterSet[chapterNumber] {
				continue
			}
			name := chapter.GetTitle()
			if name == "" {
				name = fmt.Sprintf("Chapter %s", chapter.GetChapterNum())
			} else {
				name = fmt.Sprintf("Chapter %s - %s", chapter.GetChapterNum(), name)
			}

			parsedCreatedDate, _ := time.Parse("2006-01-02T15:04:05-07:00", chapter.Attributes.PublishAt)

			var volume string
			if chapter.Attributes.Volume != nil {
				volume = fmt.Sprintf("Vol.%s", *chapter.Attributes.Volume)
			}

			scanlations := lo.Filter(chapter.Relationships, func(item mangodex.Relationship, index int) bool {
				return item.Type == mangodex.ScanlationGroupRel
			})
			scanlationNames := lo.Map(scanlations, func(item mangodex.Relationship, index int) string {
				return item.Attributes.(*mangodex.ScanlationGroupAttributes).Name
			})

			chapters = append(chapters, &source.Chapter{
				Name:        name,
				Index:       chapterIndex,
				Number:      chapterNumber,
				Scanlations: scanlationNames,
				ID:          chapter.ID,
				URL:         fmt.Sprintf("https://mangadex.org/chapter/%s", chapter.ID),
				Manga:       manga,
				Volume:      volume,
				ChapterDate: &parsedCreatedDate,
			})
			chapterIndex++
			chapterSet[chapterNumber] = true
		}
		currOffset += 500
		if currOffset >= list.Total {
			break
		}
	}

	slices.SortFunc(chapters, func(a, b *source.Chapter) bool {
		return a.Index < b.Index
	})

	chapters, err := removeDuplicate(chapters)
	if err != nil {
		return nil, err
	}

	manga.Chapters = chapters
	_ = m.cache.chapters.Set(manga.URL, chapters)
	return chapters, nil
}
