// http is sample package
package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/go-resty/resty/v2"
)

type httpClient interface {
	Request()
	Get()
}

type httpClientImpl struct{}

func (h *httpClientImpl) Get(
	baseURL string, header, query map[string]string) (
	respBody []byte, err error, statusCode int) {
	// Create a Resty Client
	client := resty.New()
	resp, err := client.R().
		SetQueryParams(query).
		SetHeaders(header).
		Get(baseURL)
	return resp.Body(), err, resp.StatusCode()
}

func (h *httpClientImpl) Request(
	endpoint, method, apipath string,
	header, query map[string]string,
	body []byte) (
	respBody []byte,
	err error,
	statusCode int,
) {
	// configure base URL
	baseURL, _ := url.Parse(endpoint)

	// join api path
	if apipath != "" {
		baseURL.Path = path.Join(baseURL.Path, apipath)
	}

	// set query
	q := baseURL.Query()
	for key, value := range query {
		q.Set(key, value)
	}
	baseURL.RawQuery = q.Encode()

	// create request URL
	reqURL := baseURL.String()

	// configure request body
	reqBody := bytes.NewReader(body)

	// create new http request
	req, err := http.NewRequest(method, reqURL, reqBody)
	if err != nil {
		return nil, err, 0
	}

	// configure header
	for key, value := range header {
		req.Header.Set(key, value)
	}

	// create and execute http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return respBody, err, statusCode
	}
	defer resp.Body.Close()

	// configure return value
	statusCode = resp.StatusCode
	respBody, _ = ioutil.ReadAll(resp.Body)
	return respBody, nil, statusCode
}

func NewHttpClientImpl() *httpClientImpl {
	return &httpClientImpl{}
}
