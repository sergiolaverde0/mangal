package mangadex

import (
	"github.com/sergiolaverde0/mangal/key"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	"testing"
)

func init() {
	viper.Set(key.MangadexLanguage, "en")
}

var mangadex = New()

func TestMangadex_Search(t *testing.T) {
	Convey("Given a mangadex instance", t, func() {
		Convey("When searching for a manga", func() {
			mangas, err := mangadex.Search("Revenge of the Sword Clan's Hound")
			Convey("Then the error should be nil", func() {
				So(err, ShouldBeNil)

				Convey("And the result should be a list of mangas", func() {
					So(len(mangas), ShouldBeGreaterThan, 0)

					Convey("And each manga should have a name, URL and ID", func() {
						for _, manga := range mangas {
							So(manga.Name, ShouldNotBeEmpty)
							So(manga.URL, ShouldNotBeEmpty)
							So(manga.ID, ShouldNotBeEmpty)
						}
					})
					Convey("When gettings chapters for the first manga", func() {
						err := mangadex.LoadChaptersOf(mangas[0])
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
										So(chapter.Scanlations, ShouldNotBeEmpty)
										So(chapter.Number, ShouldBeGreaterThanOrEqualTo, 0)
									}
								})
							})
						})
					})
				})
			})
		})
	})
}
