// http is sample package

package http

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_httpClientImpl_Request(t *testing.T) {
	// configure variables
	testMethodGet := "GET"
	testApiPath := "/apipath"
	testQuery := map[string]string{"id": "hoge"}
	testHeader := map[string]string{"Api-Key": "sample api key"}
	testBody := []byte("test body")
	expectedResponse200 := "expected response body"
	expectedResponse401Error := "401 Unauthorized"

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
			httpHandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, expectedResponse200) }),
			wantRespBody:    expectedResponse200,
			wantStatusCode:  200,
			wantErr:         "",
		},
		{
			name: "httpClientImpl.Request 401Unauthorized ReturnsEqualsError",
			h:    &httpClientImpl{},
			args: args{
				endpoint: "",
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
			wantStatusCode: 401,
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
			ts := httptest.NewServer(tt.httpHandlerFunc)
			if tt.args.endpoint == "" {
				tt.args.endpoint = ts.URL
			}
			gotRespBody, err, gotStatusCode := tt.h.Request(tt.args.endpoint, tt.args.method, tt.args.apipath, tt.args.header, tt.args.query, tt.args.body)
			ts.Close()

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
