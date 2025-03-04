package flaresolverr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	key "github.com/sergiolaverde0/mangal/key"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type TransportFlaresolevrr struct {
	client      *http.Client
	uuid        uuid.UUID
	mutex       sync.Mutex
	uuidBuilder sync.Once
}

func NewTransport() *TransportFlaresolevrr {

	return &TransportFlaresolevrr{
		client: new(http.Client),
		uuid:   uuid.Nil,
	}
}

func unmarshalJSON[T any](b []byte) (v T, err error) {
	return v, json.Unmarshal(b, &v)
}

func (t *TransportFlaresolevrr) RoundTrip(r *http.Request) (*http.Response, error) {

	t.uuidBuilder.Do(func() {
		t.uuid = uuid.New()
	})
	t.mutex.Lock()
	defer t.mutex.Unlock()

	var req []byte
	var err error
	deadline, ok := r.Context().Deadline()
	var timeout int
	if ok {
		timeout = int(deadline.Sub(time.Now()).Milliseconds()) - 1000
	} else {
		timeout = 30000
	}
	switch r.Method {
	case "GET":
		req, err = json.Marshal(request{
			Cmd:        "request.get",
			URL:        r.URL.String(),
			MaxTimeout: timeout,
			Session:    t.uuid.String(),
		})
		break
	case "POST":
		content, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		req, err = json.Marshal(request{
			Cmd:        "request.post",
			URL:        r.URL.String(),
			MaxTimeout: timeout,
			Session:    t.uuid.String(),
			PostData:   string(content),
		})
		break
	default:
		return nil, errors.New("only support GET and POST methods")
	}

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(r.Context(), "POST", viper.GetString(key.SourceHeadlessFlaresolverrURL), bytes.NewBuffer(req))
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	resp, err := t.client.Do(request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	flareResponse, err := unmarshalJSON[response](body)

	if err != nil {
		return nil, err
	}

	if flareResponse.Status == "error" {
		return nil, errors.New(flareResponse.Message)
	}

	response := &http.Response{
		StatusCode: flareResponse.Solution.Status,
		Header:     flareResponse.Solution.Headers,
		Body:       newContainer(&flareResponse),
		Request:    r,
	}
	if response.Header.Get("Content-Type") == "" {
		response.Header.Set("Content-Type", "text/html")
		response.Header.Set("Content-Length", strconv.Itoa(len(body)))
	}

	return response, nil
}

func (t *TransportFlaresolevrr) Close() error {
	if t.uuid == uuid.Nil {
		return nil
	}
	req, err := json.Marshal(request{
		Cmd:     "sessions.destroy",
		Session: t.uuid.String(),
	})
	if err != nil {
		return err
	}
	request, err := http.NewRequestWithContext(context.Background(), "POST", viper.GetString(key.SourceHeadlessFlaresolverrURL), bytes.NewBuffer(req))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	resp, err := t.client.Do(request)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil

}
