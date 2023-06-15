package headless

import (
	"github.com/metafates/mangal/key"
	"github.com/metafates/mangal/log"
	"github.com/spf13/viper"
	"net/http"
	"sync"
)

var (
	transportInstance TransportHeadless
	once              sync.Once
)

type TransportHeadless interface {
	// RoundTrip executes a single HTTP transaction, returning
	// a Response for the provided Request.
	//
	// RoundTrip should not attempt to interpret the response. In
	// particular, RoundTrip must return err == nil if it obtained
	// a response, regardless of the response's HTTP status code.
	// A non-nil err should be reserved for failure to obtain a
	// response. Similarly, RoundTrip should not attempt to
	// handle higher-level protocol details such as redirects,
	// authentication, or cookies.
	//
	// RoundTrip should not modify the request, except for
	// consuming and closing the Request's Body. RoundTrip may
	// read fields of the request in a separate goroutine. Callers
	// should not mutate or reuse the request until the Response's
	// Body has been closed.
	//
	// RoundTrip must always close the body, including on errors,
	// but depending on the implementation may do so in a separate
	// goroutine even after RoundTrip returns. This means that
	// callers wanting to reuse the body for subsequent requests
	// must arrange to wait for the Close call before doing so.
	//
	// The Request's URL and Header fields must be initialized.
	RoundTrip(r *http.Request) (*http.Response, error)

	// Close any cleanup when closing the transport
	Close() error
}

func GetTransportSingleton() TransportHeadless {
	once.Do(func() {
		if viper.GetBool(key.SourceHeadlessUseFlaresolverr) && viper.GetString(key.SourceHeadlessFlaresolverrURL) != "" {
			result, err := http.Get(viper.GetString(key.SourceHeadlessFlaresolverrURL))
			defer result.Body.Close()
			if err != nil || result.StatusCode != 200 {
				log.Error("Couldn't connect to flaresolverr, falling back to rod")
				transportInstance = newRod()
				return
			}
			transportInstance = NewFlareSolverr()
			return
		}
		transportInstance = newRod()
	})
	return transportInstance
}
