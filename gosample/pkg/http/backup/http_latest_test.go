// http is sample package

package http

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func Test_httpClientImpl_Request(t *testing.T) {
	// configure retry variables
	retryMax := 3
	retryWaitMin := 1 * time.Second
	retryWaitMax := 5 * time.Second
	httpRequestTimeout := 5 * time.Second
	// configure variables
	testMethodGet := "GET"
	testApiPath := "/apipath"
	testQuery := map[string]string{"id": "hoge"}
	testHeader := map[string]string{"Api-Key": "sample api key"}
	testBody := []byte("test body")
	expectedResponse200 := "expected response body"
	expectedResponse401Error := "401 Unauthorized"
	expectedResponse500Error := "500 Internal Server Error"
	customEndpoint := "127.0.0.1:8080"

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
			h: &httpClientImpl{
				retryMax:           retryMax,
				retryWaitMin:       retryWaitMin,
				retryWaitMax:       retryWaitMax,
				httpRequestTimeout: httpRequestTimeout,
			},
			args: args{
				endpoint: "customEndpoint", // uses custom endpoint
				method:   testMethodGet,
				apipath:  testApiPath,
				header:   testHeader,
				query:    testQuery,
				body:     testBody,
			},
			httpHandlerFunc: http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintf(w, expectedResponse200)
				},
			),
			wantRespBody:   expectedResponse200,
			wantStatusCode: http.StatusOK,
			wantErr:        "",
		},
		{
			name: "httpClientImpl.Request URLParseError ReturnsEqualsError",
			h:    &httpClientImpl{},
			args: args{
				endpoint: "127.0.0.1:8080",
				method:   testMethodGet,
				apipath:  testApiPath,
				header:   testHeader,
				query:    testQuery,
				body:     testBody,
			},
			httpHandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			wantRespBody:    "",
			wantStatusCode:  0,
			wantErr:         `parse "127.0.0.1:8080": first path segment in URL cannot contain colon`,
		},
		{
			name: "httpClientImpl.Request 401Unauthorized ReturnsEqualsError",
			h: &httpClientImpl{
				retryMax:           retryMax,
				retryWaitMin:       retryWaitMin,
				retryWaitMax:       retryWaitMax,
				httpRequestTimeout: httpRequestTimeout,
			},
			args: args{
				endpoint: "customEndpoint", // uses custom endpoint
				method:   testMethodGet,
				apipath:  testApiPath,
				header:   testHeader,
				query:    testQuery,
				body:     testBody,
			},
			httpHandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, expectedResponse401Error, http.StatusUnauthorized)
			}),
			wantRespBody:   expectedResponse401Error + "\n",
			wantStatusCode: http.StatusUnauthorized,
			wantErr:        "",
		},
		{
			name: "httpClientImpl.Request 500InternalServerError ReturnsEqualsError",
			h: &httpClientImpl{
				retryMax:           retryMax,
				retryWaitMin:       retryWaitMin,
				retryWaitMax:       retryWaitMax,
				httpRequestTimeout: httpRequestTimeout,
			},
			args: args{
				endpoint: "customEndpoint", // uses custom endpoint
				method:   testMethodGet,
				apipath:  testApiPath,
				header:   testHeader,
				query:    testQuery,
				body:     testBody,
			},
			httpHandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, expectedResponse500Error, http.StatusInternalServerError)
			}),
			wantRespBody:   "",
			wantStatusCode: 0,
			wantErr:        `Get "http://127.0.0.1:8080/apipath?id=hoge": GET http://127.0.0.1:8080/apipath?id=hoge giving up after 4 attempt(s)`,
		},
		{
			name: "httpClientImpl.Request UnsupportedProtcolIsSpecifed ReturnsEqualsError",
			h: &httpClientImpl{
				retryMax:           retryMax,
				retryWaitMin:       retryWaitMin,
				retryWaitMax:       retryWaitMax,
				httpRequestTimeout: httpRequestTimeout,
			},
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
			wantErr:         `Get "WrongURL/apipath?id=hoge": GET WrongURL/apipath?id=hoge giving up after 1 attempt(s): Get "WrongURL/apipath?id=hoge": unsupported protocol scheme ""`,
		},
		{
			name: "httpClientImpl.Request invalidResponseBody ReturnsEqualsError",
			h: &httpClientImpl{
				retryMax:           retryMax,
				retryWaitMin:       retryWaitMin,
				retryWaitMax:       retryWaitMax,
				httpRequestTimeout: httpRequestTimeout,
			},
			args: args{
				endpoint: "customEndpoint",
				method:   testMethodGet,
				apipath:  testApiPath,
				header:   testHeader,
				query:    testQuery,
				body:     testBody,
			},
			httpHandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Header().Set("Content-Length", "1") }),
			wantRespBody:    "",
			wantStatusCode:  200,
			wantErr:         io.ErrUnexpectedEOF.Error(),
		},
	}

	isErrStringEqualsWantErrString := func(err error, wantErrString string) bool {
		errString := ""
		if err != nil {
			errString = err.Error()
		}
		if errString == wantErrString {
			return true
		}
		return false
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ==================================================
			// create mock server
			// ==================================================
			l, err := net.Listen("tcp", customEndpoint)
			if err != nil {
				log.Fatalln("Unit Test Configuration Error: HTTP Mock Server:", err)
			}
			ts := httptest.NewUnstartedServer(tt.httpHandlerFunc)
			ts.Listener.Close()
			ts.Listener = l
			ts.Start()
			// ==================================================
			// run unit test
			// ==================================================
			if tt.args.endpoint == "customEndpoint" {
				tt.args.endpoint = ts.URL
			}
			gotRespBody, err, gotStatusCode := tt.h.Request(tt.args.endpoint, tt.args.method, tt.args.apipath, tt.args.header, tt.args.query, tt.args.body)
			// ==================================================
			// close server
			// ==================================================
			ts.Close()
			// ==================================================
			// validation
			// ==================================================
			if !isErrStringEqualsWantErrString(err, tt.wantErr) {
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
	retryMax := 3
	retryWaitMin := 1 * time.Second
	retryWaitMax := 5 * time.Second
	httpRequestTimeout := 5 * time.Second
	type args struct {
		retryMax           int
		retryWaitMin       time.Duration
		retryWaitMax       time.Duration
		httpRequestTimeout time.Duration
	}
	tests := []struct {
		name string
		args args
		want *httpClientImpl
	}{
		{
			name: "NewHttpClientImpl CreateInstance ReturnsEqualsInstance",
			args: args{
				retryMax:           retryMax,
				retryWaitMin:       retryWaitMin,
				retryWaitMax:       retryWaitMax,
				httpRequestTimeout: httpRequestTimeout,
			},
			want: &httpClientImpl{
				retryMax:           retryMax,
				retryWaitMin:       retryWaitMin,
				retryWaitMax:       retryWaitMax,
				httpRequestTimeout: httpRequestTimeout,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHttpClientImpl(tt.args.retryMax, tt.args.retryWaitMin, tt.args.retryWaitMax, tt.args.httpRequestTimeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHttpClientImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}
