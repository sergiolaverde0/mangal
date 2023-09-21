package asurascans

import (
	"github.com/belphemur/mangal/key"
	"github.com/belphemur/mangal/provider/generic"
	"github.com/belphemur/mangal/provider/generic/headless"
	"github.com/belphemur/mangal/util"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	"testing"
)

func init() {
	viper.Set(key.SourceHeadlessUseFlaresolverr, false)
	viper.Set(key.SourceHeadlessFlaresolverrURL, "http://localhost:8191/v1")
}

func TestAsurascans(t *testing.T) {
	util.SkipCI(t)
	defer headless.GetTransportSingleton().Close()
	Convey("Given a asurascans instance", t, func() {
		asurascans := generic.New(Config)
		Convey("When searching for a manga", func() {
			mangas, err := asurascans.Search("Boundless")
			Convey("Then the error should be nil", func() {
				So(err, ShouldBeNil)

				Convey("And the result should be a list of mangas", func() {
					So(len(mangas), ShouldBeGreaterThan, 0)

					Convey("And each manga should have a name and URL", func() {
						for _, manga := range mangas {
							So(manga.Name, ShouldNotBeEmpty)
							So(manga.URL, ShouldNotBeEmpty)
						}
					})

					Convey("When gettings chapters for the first manga", func() {
						err := asurascans.LoadChaptersOf(mangas[0])
						chapters := mangas[0].Chapters
						Convey("Then the error should be nil", func() {
							So(err, ShouldBeNil)

							Convey("And the result should be a list of chapters", func() {
								So(len(chapters), ShouldBeGreaterThan, 0)

								Convey("And each chapter should have a name, URL and manga relation", func() {
									for _, chapter := range chapters {
										So(chapter.Name, ShouldNotBeEmpty)
										So(chapter.URL, ShouldNotBeEmpty)
										So(chapter.Manga, ShouldEqual, mangas[0])
										So(chapter.ChapterDate, ShouldNotBeEmpty)
										So(chapter.Number, ShouldBeGreaterThanOrEqualTo, 0)
									}
								})

								Convey("When getting pages for the first chapter", func() {
									err := asurascans.LoadPagesOf(chapters[0])
									pages := chapters[0].Pages
									Convey("Then the error should be nil", func() {
										So(err, ShouldBeNil)

										Convey("And the result should be a list of pages", func() {
											So(len(pages), ShouldBeGreaterThan, 0)

											Convey("And each page should have a URL, non nil contents and chapter relation", func() {
												for _, page := range pages {
													So(page.URL, ShouldNotBeEmpty)
													So(page.Chapter, ShouldEqual, chapters[0])
												}
											})
										})
									})
								})
							})
						})
					})
				})
			})
		})
	})
}
