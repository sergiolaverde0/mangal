package flaresolverr

import (
	"context"
	"github.com/belphemur/mangal/key"
	"github.com/belphemur/mangal/util"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	"net/http"
	"testing"
	"time"
)

func init() {
	viper.Set(key.SourceHeadlessUseFlaresolverr, true)
	viper.Set(key.SourceHeadlessFlaresolverrURL, "http://localhost:8191/v1")
}

func TestTransportFlaresolevrr(t *testing.T) {
	if util.SkipCI(t) {
		return
	}

	transport := NewTransport()

	client := http.Client{
		Transport: transport,
	}
	defer client.CloseIdleConnections()
	Convey("Flaresolverr can go on ReaperScans", t, func() {
		ctx, _ := context.WithTimeout(context.Background(), 120*time.Second)
		req, err := http.NewRequestWithContext(ctx, "GET", "https://www.reaperscans.com", nil)
		So(err, ShouldBeNil)

		resp, err := client.Do(req)
		So(err, ShouldBeNil)
		defer resp.Body.Close()
		So(resp.StatusCode, ShouldEqual, 200)

		Convey("Flaresolverr session should be destroyed when finished", func() {
			err := transport.Close()
			So(err, ShouldBeNil)
		})
	})

}
