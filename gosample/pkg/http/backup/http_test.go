// http is sample package

package http_old

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// URL Creator
// urlCreator := func(endpoint, apipath string, query map[string]string) string {
// 	// configure base URL
// 	baseURL, err := url.Parse("http://" + endpoint)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// join api path
// 	if apipath != "" {
// 		baseURL.Path = path.Join(baseURL.Path, apipath)
// 	}
// 	// set query
// 	q := baseURL.Query()
// 	for key, value := range query {
// 		q.Set(key, value)
// 	}
// 	baseURL.RawQuery = q.Encode()
// 	return baseURL.String()
// }
// fmt.Println(urlCreator(customEndpoint, testApiPath, testQuery))

// wantErr: fmt.Errorf(
// 	`Get "%s": %s %s giving up after %d attempt(s)`,
// 	urlCreator(customEndpoint, testApiPath, testQuery), testMethodGet, urlCreator(customEndpoint, testApiPath, testQuery), retryMax+1).Error(),

// func Test_httpClientImpl_Get(t *testing.T) {
// 	// configure variables
// 	testQuery := map[string]string{
// 		"id": "hoge",
// 	}
// 	testHeader := map[string]string{
// 		"Api-Key": "sample api key",
// 	}
// 	type args struct {
// 		baseURL string
// 		header  map[string]string
// 		query   map[string]string
// 	}
// 	tests := []struct {
// 		name            string
// 		h               *httpClientImpl
// 		args            args
// 		httpHandlerFunc http.HandlerFunc
// 		wantRespBody    string
// 		wantStatusCode  int
// 		wantErr         bool
// 		testSetup       func()
// 		testTeardown    func()
// 	}{
// 		{
// 			name: "httpClinetImpl.Get GetRequest ReturnsErrorEqualsNil",
// 			h:    &httpClientImpl{},
// 			args: args{
// 				baseURL: "",
// 				header:  testHeader,
// 				query:   testQuery,
// 			},
// 			httpHandlerFunc: http.HandlerFunc(
// 				func(w http.ResponseWriter, r *http.Request) {
// 					fmt.Fprintf(w, "Hello HTTP Test")
// 				},
// 			),
// 			wantRespBody:   "Hello HTTP Test",
// 			wantStatusCode: 200,
// 			wantErr:        false,
// 		},
// 		{
// 			name: "httpClinetImpl.Get 404NotFound ReturnsErrorEqualsError",
// 			h:    &httpClientImpl{},
// 			args: args{
// 				baseURL: "",
// 				header:  testHeader,
// 				query:   testQuery,
// 			},
// 			httpHandlerFunc: http.HandlerFunc(
// 				func(w http.ResponseWriter, r *http.Request) {
// 					http.Error(w, "Not Found.", http.StatusNotFound)
// 				},
// 			),
// 			wantRespBody:   "Not Found.\n",
// 			wantStatusCode: http.StatusNotFound,
// 			wantErr:        false,
// 		},
// 		{
// 			name: "httpClinetImpl.Get RequestError ReturnsErrorEqualsError",
// 			h:    &httpClientImpl{},
// 			args: args{
// 				baseURL: "wrong url",
// 				header:  testHeader,
// 				query:   testQuery,
// 			},
// 			httpHandlerFunc: http.HandlerFunc(
// 				func(w http.ResponseWriter, r *http.Request) {},
// 			),
// 			wantRespBody:   "",
// 			wantStatusCode: 0,
// 			wantErr:        true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ts := httptest.NewServer(tt.httpHandlerFunc)
// 			var baseURL string
// 			if tt.args.baseURL != "" {
// 				baseURL = tt.args.baseURL
// 			} else {
// 				baseURL = ts.URL
// 			}
// 			gotRespBody, err, gotStatusCode := tt.h.Get(baseURL, tt.args.header, tt.args.query)
// 			ts.Close()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("httpClientImpl.Get() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if gotRespBody != tt.wantRespBody {
// 				t.Errorf("httpClientImpl.Get() gotRespBody = %v, want %v", gotRespBody, tt.wantRespBody)
// 			}
// 			if gotStatusCode != tt.wantStatusCode {
// 				t.Errorf("httpClientImpl.Get() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
// 			}
// 		})
// 	}
// }

func badStringError(what, val string) error { return fmt.Errorf("%s %q", what, val) }

func Test_httpClientImpl_Request(t *testing.T) {
	// configure variables
	testMethodGet := "GET"
	testApiPath := "/apipath"
	testQuery := map[string]string{"id": "hoge"}
	testHeader := map[string]string{"Api-Key": "sample api key"}
	testBody := []byte("test body")
	expectedResponseBody := "expected response body"
	type args struct {
		endpoint string
		method   string
		apipath  string
		header   map[string]string
		query    map[string]string
		body     []byte
	}
	tests := []struct {
		name            string
		h               *httpClientImpl
		args            args
		httpHandlerFunc http.HandlerFunc
		wantRespBody    string
		wantStatusCode  int
		wantErr         string
	}{
		{
			name: "httpClientImpl.Request GetRequst ReturnsEqualsNil",
			h:    &httpClientImpl{},
			args: args{
				endpoint: "",
				method:   testMethodGet,
				apipath:  testApiPath,
				header:   testHeader,
				query:    testQuery,
				body:     testBody,
			},
			httpHandlerFunc: http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintf(w, expectedResponseBody)
				},
			),
			wantRespBody:   expectedResponseBody,
			wantStatusCode: 200,
			wantErr:        "",
		},
		{
			name: "httpClientImpl.Request UnsupportedProtcolIsSpecifed ReturnsEqualsError",
			h:    &httpClientImpl{},
			args: args{
				endpoint: "WrongURL",
				method:   testMethodGet,
				apipath:  testApiPath,
				header:   testHeader,
				query:    testQuery,
				body:     testBody,
			},
			httpHandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			wantRespBody:    "",
			wantStatusCode:  0,
			wantErr:         `Get "WrongURL/apipath?id=hoge": unsupported protocol scheme ""`,
		},
		{
			name: "httpClientImpl.Request invalidResponseBody ReturnsEqualsError",
			h:    &httpClientImpl{},
			args: args{
				endpoint: "",
				method:   testMethodGet,
				apipath:  testApiPath,
				header:   testHeader,
				query:    testQuery,
				body:     testBody,
			},
			httpHandlerFunc: http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Length", "1")
				},
			),
			wantRespBody:   "",
			wantStatusCode: 200,
			wantErr:        io.ErrUnexpectedEOF.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(tt.httpHandlerFunc)
			var endpoint string
			if tt.args.endpoint == "" {
				endpoint = ts.URL
			} else {
				endpoint = tt.args.endpoint
			}
			gotRespBody, err, gotStatusCode := tt.h.Request(endpoint, tt.args.method, tt.args.apipath, tt.args.header, tt.args.query, tt.args.body)

			var errString string
			if err != nil {
				errString = err.Error()
			}
			if errString != tt.wantErr {
				t.Errorf("httpClientImpl.Request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRespBody != tt.wantRespBody {
				t.Errorf("httpClientImpl.Request() gotRespBody = %v, want %v", gotRespBody, tt.wantRespBody)
			}
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("httpClientImpl.Request() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
		})
	}
}

func TestNewHttpClientImpl(t *testing.T) {
	tests := []struct {
		name string
		want *httpClientImpl
	}{
		{
			name: "NewHttpClientImpl CreateInstance ReturnsEqualsInstance",
			want: &httpClientImpl{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHttpClientImpl(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHttpClientImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}
