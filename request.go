package freelancer

import "net/url"

type params map[string]interface{}

type request struct {
	method   string
	endpoint string
	query    url.Values
	form     url.Values
}
