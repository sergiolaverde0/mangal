package flaresolverr

import (
	"io"
	"strings"
)

type response struct {
	Solution struct {
		URL      string              `json:"url"`
		Status   int                 `json:"status"`
		Headers  map[string][]string `json:"headers"`
		Response string              `json:"response"`
		Cookies  []struct {
			Name     string  `json:"name"`
			Value    string  `json:"value"`
			Domain   string  `json:"domain"`
			Path     string  `json:"path"`
			Expires  float64 `json:"expires"`
			Size     int     `json:"size"`
			HTTPOnly bool    `json:"httpOnly"`
			Secure   bool    `json:"secure"`
			Session  bool    `json:"session"`
			SameSite string  `json:"sameSite"`
		} `json:"cookies"`
		UserAgent string `json:"userAgent"`
	} `json:"solution"`
	Status         string `json:"status"`
	Message        string `json:"message"`
	StartTimestamp int64  `json:"startTimestamp"`
	EndTimestamp   int64  `json:"endTimestamp"`
	Version        string `json:"version"`
}

type container struct {
	io.Reader
}

func newContainer(resp *response) *container {
	return &container{Reader: strings.NewReader(resp.Solution.Response)}
}

func (c container) Read(p []byte) (n int, err error) {
	return c.Reader.Read(p)
}

func (c container) Close() error {
	return nil
}
