package source

import (
	"fmt"
	"github.com/sergiolaverde0/mangal/constant"
	"github.com/sergiolaverde0/mangal/filesystem"
	"github.com/sergiolaverde0/mangal/key"
	"github.com/sergiolaverde0/mangal/log"
	"github.com/sergiolaverde0/mangal/style"
	"github.com/sergiolaverde0/mangal/util"
	"github.com/dustin/go-humanize"
	"github.com/samber/lo"
	"github.com/samber/mo"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// Chapter is a struct that represents a chapter of a manga.
type Chapter struct {
	// Name of the chapter
	Name string `json:"name" jsonschema:"description=Name of the chapter"`
	// Number represent the chapter number like 145.5 0.5
	Number float32 `json:"number" jsonschema:"description=Number of the chapter as defined in its source"`
	// URL of the chapter
	URL string `json:"url" jsonschema:"description=URL of the chapter"`
	// Index of the chapter in the manga.
	Index uint16 `json:"index" jsonschema:"description=Index of the chapter in the manga"`
	// ID of the chapter in the source.
	ID string `json:"id" jsonschema:"description=ID of the chapter in the source"`
	// Volume which the chapter belongs to.
	Volume string `json:"volume" jsonschema:"description=Volume which the chapter belongs to"`
	// Manga that the chapter belongs to.
	Manga *Manga `json:"-"`
	// Pages of the chapter.
	Pages []*Page `json:"pages" jsonschema:"description=Pages of the chapter"`
	// Date when the chapter was released.
	ChapterDate *time.Time `json:"chapter_date" jsonschema:"description=Date when the chapter was released"`
	// Scanlations team that translated the chapter.
	Scanlations []string `json:"scanlation,omitempty" jsonschema:"description=Scanlations team(s) that translated the chapter"`

	isDownloaded mo.Option[bool]
	size         uint64
}

func (c *Chapter) String() string {
	return c.Name
}

func (c *Chapter) UpdatePages(pages []*Page) {
	c.Pages = pages
	c.size = lo.SumBy(pages, func(item *Page) uint64 {
		return item.Size
	})
}

// DownloadPages downloads the Pages contents of the Chapter.
// Pages needs to be set before calling this function.
func (c *Chapter) DownloadPages(temp bool, progress func(string)) (err error) {
	if c.Pages == nil || len(c.Pages) == 0 {
		log.Error("Couldn't find any pages")
		return fmt.Errorf("couldn't find pages to download for chapter %d", c.Index)
	}

	c.size = 0
	status := func() string {
		return fmt.Sprintf(
			"Downloading %s %s",
			util.Quantify(len(c.Pages), "page", "pages"),
			style.Faint(c.SizeHuman()),
		)
	}

	progress(status())
	wg := sync.WaitGroup{}
	wg.Add(len(c.Pages))

	for _, page := range c.Pages {
		if page == nil {
			return fmt.Errorf("page #%d is empty, aborting download", page.Index)
		}

		d := func(page *Page) {
			defer wg.Done()

			// if at any point, an error is encountered, stop downloading other pages
			if err != nil {
				return
			}

			err = page.Download()
			c.size += page.Size
			progress(status())
		}

		if viper.GetBool(key.DownloaderAsync) {
			go d(page)
		} else {
			d(page)
		}
	}

	wg.Wait()

	if err != nil {
		c.isDownloaded = mo.Some(false)
		return err
	}

	c.isDownloaded = mo.Some(!temp)
	return
}

// formattedName of the chapter according to the template in the config.
func (c *Chapter) formattedName() (name string) {
	name = viper.GetString(key.DownloaderChapterNameTemplate)

	var sourceName string
	if c.Source() != nil {
		sourceName = c.Source().Name()
	}

	for variable, value := range map[string]string{
		"manga":          c.Manga.Name,
		"chapter":        c.Name,
		"number":         strings.TrimSuffix(fmt.Sprintf("%.1f", c.Number), ".0"),
		"padded-number":  strings.TrimSuffix(fmt.Sprintf("%04.1f", c.Number), ".0"),
		"index":          fmt.Sprintf("%d", c.Index),
		"padded-index":   fmt.Sprintf("%04d", c.Index),
		"chapters-count": fmt.Sprintf("%d", len(c.Manga.Chapters)),
		"volume":         c.Volume,
		"source":         sourceName,
	} {
		name = strings.ReplaceAll(name, fmt.Sprintf("{%s}", variable), value)
	}

	return
}

// SizeHuman is the same as Size but returns a human-readable string.
func (c *Chapter) SizeHuman() string {
	return humanize.Bytes(c.size)
}

func (c *Chapter) Filename() (filename string) {
	filename = util.SanitizeFilename(c.formattedName())

	// plain format assumes that chapter is a directory with images
	// rather than a single file. So no need to add extension to it
	if f := viper.GetString(key.FormatsUse); f != constant.FormatPlain {
		return filename + "." + f
	}

	return
}

func (c *Chapter) IsDownloaded() bool {
	if c.isDownloaded.IsPresent() {
		return c.isDownloaded.MustGet()
	}

	path, _ := c.path(c.Manga.peekPath(), false)
	exists, _ := filesystem.Api().Exists(path)
	c.isDownloaded = mo.Some(exists)
	return exists
}

func (c *Chapter) path(relativeTo string, createVolumeDir bool) (path string, err error) {
	if createVolumeDir {
		path = filepath.Join(path, util.SanitizeFilename(c.Volume))
		err = filesystem.Api().MkdirAll(path, os.ModePerm)
		if err != nil {
			return
		}
	}

	path = filepath.Join(relativeTo, c.Filename())
	return
}

func (c *Chapter) Path(temp bool) (path string, err error) {
	var manga string
	manga, err = c.Manga.Path(temp)
	if err != nil {
		return
	}

	return c.path(manga, c.Volume != "" && viper.GetBool(key.DownloaderCreateVolumeDir))
}

func (c *Chapter) Source() Source {
	return c.Manga.Source
}

func (c *Chapter) ComicInfo() *ComicInfo {
	var day, month, year = getChapterDate(c)

	return &ComicInfo{
		XmlnsXsd: "http://www.w3.org/2001/XMLSchema",
		XmlnsXsi: "http://www.w3.org/2001/XMLSchema-instance",

		Title:      c.Name,
		Series:     c.Manga.Name,
		Number:     strings.TrimSuffix(fmt.Sprintf("%.1f", c.Number), ".0"),
		Web:        c.URL,
		Genre:      strings.Join(c.Manga.Metadata.Genres, ","),
		PageCount:  len(c.Pages),
		Summary:    c.Manga.Metadata.Summary,
		Count:      c.Manga.Metadata.Chapters,
		Characters: strings.Join(c.Manga.Metadata.Characters, ","),
		Year:       year,
		Month:      month,
		Day:        day,
		Writer:     strings.Join(c.Manga.Metadata.Staff.Story, ","),
		Penciller:  strings.Join(c.Manga.Metadata.Staff.Art, ","),
		Letterer:   strings.Join(c.Manga.Metadata.Staff.Lettering, ","),
		Translator: strings.Join(c.Manga.Metadata.Staff.Translation, ","),
		Tags:       strings.Join(c.Manga.Metadata.Tags, ","),
		Notes:      "Downloaded with Mangal. https://github.com/metafates/mangal",
		Teams:      strings.Join(c.Scanlations, ", "),
		Manga:      "YesAndRightToLeft",
	}
}

func getChapterDate(c *Chapter) (int, int, int) {
	var (
		day, month, year int
	)
	if !viper.GetBool(key.MetadataComicInfoXMLAddDate) {
		return day, month, year

	}
	if c.ChapterDate != nil {
		day = c.ChapterDate.Day()
		month = int(c.ChapterDate.Month())
		year = c.ChapterDate.Year()
		return day, month, year
	}

	if viper.GetBool(key.MetadataComicInfoXMLAlternativeDate) {
		// get current date
		t := time.Now()
		day = t.Day()
		month = int(t.Month())
		year = t.Year()
	} else {
		day = c.Manga.Metadata.StartDate.Day
		month = c.Manga.Metadata.StartDate.Month
		year = c.Manga.Metadata.StartDate.Year
	}

	return day, month, year
}
