package anilist

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSearch(t *testing.T) {
	Convey(`Given a query "Death Note"`, t, func() {
		query := "Death Note"
		Convey(`When I search for it`, func() {
			results, err := SearchByName(query)
			Convey(`Then I should get a result`, func() {
				So(err, ShouldBeNil)
				So(results, ShouldNotBeEmpty)
				Convey(`And the first result should be "Death Note"`, func() {
					So(results[0].Title.English, ShouldEqual, "Death Note")
				})
			})
		})
	})
}
func TestSearchMangaUpdates(t *testing.T) {
	Convey(`Given a query "Death Note"`, t, func() {
		query := "Insanely-Talented Player"
		Convey(`When I search for it`, func() {
			results, err := SearchByName(query)
			Convey(`Then I should get a result`, func() {
				So(err, ShouldBeNil)
				So(results, ShouldNotBeEmpty)
				Convey(`And the first result should be "Insanely-Talented Player"`, func() {
					So(results[0].Title.English, ShouldEqual, "Insanely Talented Player")
				})
			})
		})
	})
}
