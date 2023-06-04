package asurascans

import (
	"github.com/metafates/mangal/provider/generic"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAsurascans(t *testing.T) {
	Convey("Given a asurascans instance", t, func() {
		asurascans := generic.New(Config)
		Convey("When searching for a manga", func() {
			mangas, err := asurascans.Search("Solo")
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
						chapters, err := asurascans.ChaptersOf(mangas[0])
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
									pages, err := asurascans.PagesOf(chapters[0])
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