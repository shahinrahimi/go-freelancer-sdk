package freelancer

import (
	"io"
	"net/http"
	"net/url"
)

type params map[string]interface{}

type request struct {
	method   string
	endpoint string
	query    url.Values
	form     url.Values
	header   http.Header
	body     io.Reader
	fullURL  string
}
