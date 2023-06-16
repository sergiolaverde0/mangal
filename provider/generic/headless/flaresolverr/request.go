package flaresolverr

type request struct {
	Cmd        string `json:"cmd"`
	URL        string `json:"url,omitempty"`
	MaxTimeout int    `json:"maxTimeout,omitempty"`
	Session    string `json:"session,omitempty"`
	PostData   string `json:"postData,omitempty"`
}
