package headless

import (
	"github.com/metafates/mangal/key"
	"github.com/metafates/mangal/log"
	"github.com/metafates/mangal/provider/generic/headless/flaresolverr"
	"github.com/metafates/mangal/provider/generic/headless/rod"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
	"sync"
)

var (
	transportInstance TransportHeadless
	once              sync.Once
)

type TransportHeadless interface {
	http.RoundTripper
	io.Closer
}

func GetTransportSingleton() TransportHeadless {
	once.Do(func() {
		if viper.GetBool(key.SourceHeadlessUseFlaresolverr) && viper.GetString(key.SourceHeadlessFlaresolverrURL) != "" {
			url, err := url.Parse(viper.GetString(key.SourceHeadlessFlaresolverrURL))
			if err != nil {
				log.Error("Couldn't parse flaresolverr url, falling back to rod")
				transportInstance = rod.NewTransport()
				return
			}

			url.Path = ""
			result, err := http.Get(url.String())
			defer func() {
				if result != nil && result.Body != nil {
					result.Body.Close()
				}
			}()
			if err != nil || result.StatusCode != 200 {
				log.Error("Couldn't connect to flaresolverr, falling back to rod")
				transportInstance = rod.NewTransport()
				return
			}
			transportInstance = flaresolverr.NewTransport()
			return
		}
		transportInstance = rod.NewTransport()
	})
	return transportInstance
}
