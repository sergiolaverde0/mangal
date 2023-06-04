package headless

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"net/http"
	"runtime"
)

type Transport struct {
	browser *rod.Browser
}

func New() *Transport {
	u := launcher.New().Leakless(runtime.GOOS == "linux").MustLaunch()
	return &Transport{
		browser: rod.New().ControlURL(u).MustConnect(),
	}
}

func (t Transport) Close() error {
	return t.browser.Close()
}

func (t Transport) RoundTrip(request *http.Request) (*http.Response, error) {
	page, err := t.browser.Page(proto.TargetCreateTarget{URL: request.URL.String()})
	err = page.WaitLoad()
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return &http.Response{Body: newPageReader(page),
		StatusCode: 200,
		Header:     map[string][]string{"Content-Type": {"text/html"}},
	}, nil
}
