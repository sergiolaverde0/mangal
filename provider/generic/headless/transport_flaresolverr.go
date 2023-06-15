package headless

import (
	"bytes"
	"github.com/fahimbagar/go-flaresolverr"
	"github.com/metafates/mangal/key"
	"github.com/spf13/viper"
	"io"
	"net/http"
)

type TransportFlaresolevrr struct {
	client flaresolverr.Client
}

func NewFlareSolverr() *TransportFlaresolevrr {
	client, _ := flaresolverr.NewClient(flaresolverr.Config{
		BaseURL: viper.GetString(key.SourceHeadlessFlaresolverrURL),
	})

	return &TransportFlaresolevrr{
		client: *client,
	}
}

type flareSolverrResponse struct {
	reader io.Reader
}

func (p flareSolverrResponse) Close() error {
	return nil
}

func (p flareSolverrResponse) Read(buffer []byte) (n int, err error) {
	return p.reader.Read(buffer)
}

func (t TransportFlaresolevrr) RoundTrip(r *http.Request) (*http.Response, error) {
	var result []byte = nil
	var err error = nil

	if r.Method == "POST" {
		result, err = t.client.Post(flaresolverr.PostParams{
			URL:               r.URL.String(),
			PostData:          r.PostForm,
			MaxTimeout:        10000,
			ReturnOnlyCookies: false,
		})
	} else {
		result, err = t.client.Get(flaresolverr.GetParams{
			URL:               r.URL.String(),
			MaxTimeout:        10000,
			ReturnOnlyCookies: false,
		})
	}

	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: 200,
		Header:     map[string][]string{"Content-Type": {"text/html"}},
		Body:       &flareSolverrResponse{bytes.NewReader(result)},
	}, nil
}

func (t TransportFlaresolevrr) Close() error {
	return nil
}
