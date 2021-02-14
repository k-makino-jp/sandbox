// http is sample package
package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type httpClient interface {
	Request(
		endpoint, method, apipath string,
		header, query map[string]string,
		body []byte) (string, error, int)
}

type httpClientImpl struct {
	retryMax           int
	retryWaitMin       time.Duration
	retryWaitMax       time.Duration
	httpRequestTimeout time.Duration
}

func (h *httpClientImpl) Request(
	endpoint, method, apipath string,
	header, query map[string]string,
	body []byte) (
	string,
	error,
	int,
) {
	// configure base URL
	baseURL, err := url.Parse(endpoint)
	if err != nil {
		return "", err, 0
	}

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
		return "", err, 0
	}

	// configure header
	for key, value := range header {
		req.Header.Set(key, value)
	}

	// exponential backoff
	retryClient := retryablehttp.NewClient()
	retryClient.Logger = nil
	retryClient.RetryMax = h.retryMax
	retryClient.RetryWaitMin = h.retryWaitMin
	retryClient.RetryWaitMax = h.retryWaitMax
	retryClient.HTTPClient.Timeout = h.httpRequestTimeout
	client := retryClient.StandardClient() // *http.Client

	resp, err := client.Do(req)
	if err != nil {
		return "", err, 0
	}
	defer resp.Body.Close()

	// configure return value
	statusCode := resp.StatusCode
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err, statusCode
	}
	return string(respBodyBytes), nil, statusCode
}

func NewHttpClientImpl(retryMax int, retryWaitMin, retryWaitMax, httpRequestTimeout time.Duration) *httpClientImpl {
	return &httpClientImpl{
		retryMax:           retryMax,
		retryWaitMin:       retryWaitMin,
		retryWaitMax:       retryWaitMax,
		httpRequestTimeout: httpRequestTimeout,
	}
}
