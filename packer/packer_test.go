package packer

import (
	"github.com/metafates/mangal/constant"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGet(t *testing.T) {
	Convey("When trying to get a valid packer", t, func() {
		_, err := Get(constant.FormatCBZ)
		Convey("Then no error should be returned", func() {
			So(err, ShouldBeNil)
		})
	})

	Convey("When trying to get an invalid packer", t, func() {
		_, err := Get("kek")
		Convey("Then an error should be returned", func() {
			So(err, ShouldNotBeNil)
		})
	})
}

func TestAvailable(t *testing.T) {
	Convey("When getting the available packers", t, func() {
		packers := Available()
		Convey("Then the available packers should be returned", func() {
			So(packers, ShouldNotBeNil)
			So(len(packers), ShouldEqual, 4)
		})
	})
}
