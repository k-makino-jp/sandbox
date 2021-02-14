// http is sample package
package http_old

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

type httpClient interface {
	Get()
	Request()
}

type httpClientImpl struct {
}

// func (h *httpClientImpl) Get(
// 	baseURL string, header, query map[string]string) (
// 	respBody string, err error, statusCode int) {
// 	// Create a Resty Client
// 	client := resty.New()
// 	resp, err := client.
// 		SetRetryCount(2).
// 		SetRetryWaitTime(1 * time.Second).
// 		SetRetryMaxWaitTime(5 * time.Second).
// 		R().
// 		SetQueryParams(query).
// 		SetHeaders(header).
// 		Get(baseURL)
// 	fmt.Println(err)
// 	return string(resp.Body()), err, resp.StatusCode()
// }

func (h *httpClientImpl) Request(
	endpoint, method, apipath string,
	header, query map[string]string,
	body []byte) (
	string,
	error,
	int,
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
		return "", err, 0
	}

	// configure header
	for key, value := range header {
		req.Header.Set(key, value)
	}

	// create and execute http client
	client := &http.Client{}
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

func NewHttpClientImpl() *httpClientImpl {
	return &httpClientImpl{}
}
