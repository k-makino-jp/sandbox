// http is an package of HTTP request
package http

import (
	"github.com/go-resty/resty/v2"
)

// Get executes HTTP request with GET method.
func Get(queries map[string]string, headers map[string]string, url string) (*resty.Response, error) {
	client := resty.New()
	return client.R().
		SetQueryParams(queries).
		SetHeaders(headers).
		Get(url)
}

// Put executes HTTP request with GET method.
func Put(queries map[string]string, headers map[string]string, body string, url string) (*resty.Response, error) {
	client := resty.New()
	return client.R().
		SetQueryParams(queries).
		SetHeaders(headers).
		SetBody(body).
		Put(url)
}

// Post executes HTTP request with POST method.
func Post(queries map[string]string, headers map[string]string, body string, url string) (*resty.Response, error) {
	client := resty.New()
	return client.R().
		SetQueryParams(queries).
		SetHeaders(headers).
		SetBody(body).
		Post(url)
}
