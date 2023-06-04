package headless

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/metafates/mangal/key"
	"github.com/spf13/viper"
	"github.com/ysmood/gson"
	"net/http"
	"runtime"
	"strings"
	"time"
)

type Transport struct {
	browser *rod.Browser
}

func New() *Transport {
	u := launcher.New().Leakless(runtime.GOOS == "linux").MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()
	return &Transport{
		browser: browser,
	}
}

func (t Transport) Close() error {
	return t.browser.Close()
}

// SetExtraHeaders whether to always send extra HTTP headers with the requests from this page.
func setExtraHeaders(p *rod.Page, headers http.Header) (func(), error) {
	gsonHeaders := proto.NetworkHeaders{}

	for key, values := range headers {
		gsonHeaders[key] = gson.New(strings.Join(values, ";"))
	}

	return p.EnableDomain(&proto.NetworkEnable{}), proto.NetworkSetExtraHTTPHeaders{Headers: gsonHeaders}.Call(p)
}

func (t Transport) RoundTrip(request *http.Request) (*http.Response, error) {
	page, err := t.browser.Context(request.Context()).Page(proto.TargetCreateTarget{URL: request.URL.String()})

	if err != nil {
		return nil, err
	}

	_, err = page.SetExtraHeaders([]string{"Referer", request.Header.Get("Referer")})
	if err != nil {
		return nil, err
	}

	err = page.Navigate(request.URL.String())
	if err != nil {
		return nil, err
	}
	waitForPage := viper.GetBool(key.DownloaderWaitPageLoad)
	if waitForPage {
		err := page.WaitLoad()
		if err != nil {
			return nil, err
		}

	} else {
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return nil, err
	}
	return &http.Response{Body: newPageReader(page),
		StatusCode: 200,
		Header:     map[string][]string{"Content-Type": {"text/html"}},
	}, nil
}
