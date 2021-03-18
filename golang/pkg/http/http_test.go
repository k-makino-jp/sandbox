// http is sample package

package http

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_httpClientImpl_Request(t *testing.T) {
	// variable for http stub
	var responseCounter uint8
	customURL := "127.0.0.1:8080"
	// sample variables
	testMethodGet := "GET"
	testApiPath := "/apipath"
	testQuery := map[string]string{"id": "hoge"}
	testHeader := map[string]string{"Api-Key": "sample api key"}
	testBody := []byte("test body")
	expectedResponseBody := "expected response body"
	// configure tests
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
			name: "httpRequest GetRequst ReturnsEqualsNil",
			h:    &httpClientImpl{},
			args: args{
				endpoint: customURL,
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
			name: "httpRequest GetRequstWith1Retry ReturnsEqualsNil",
			h: &httpClientImpl{
				retryMax: 1,
				// retryWaitMin:       0,
				// retryWaitMax:       0,
				// httpRequestTimeout: 0,
			},
			args: args{
				endpoint: customURL,
				method:   testMethodGet,
				apipath:  testApiPath,
				header:   testHeader,
				query:    testQuery,
				body:     testBody,
			},
			httpHandlerFunc: http.HandlerFunc(
				// configure http stub response
				func(w http.ResponseWriter, r *http.Request) {
					responseInternalServerError := []func(w http.ResponseWriter, r *http.Request){
						// 1st response
						func(w http.ResponseWriter, r *http.Request) {
							http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
						},
						// 2nd response
						func(w http.ResponseWriter, r *http.Request) {
							fmt.Fprintf(w, expectedResponseBody)
						},
						// over 3 tries
						func(w http.ResponseWriter, r *http.Request) {
							t.Errorf("Unit Test Configuration Error: Too many retry called.")
						},
					}
					responseInternalServerError[responseCounter](w, r)
					responseCounter++
				}),
			wantRespBody:   expectedResponseBody,
			wantStatusCode: 200,
			wantErr:        "",
		},
		{
			name: "httpRequest AllGetRequestFailed ReturnsEqualsNil",
			h: &httpClientImpl{
				retryMax: 1,
				// retryWaitMin:       0,
				// retryWaitMax:       0,
				// httpRequestTimeout: 0,
			},
			args: args{
				endpoint: customURL,
				method:   testMethodGet,
				apipath:  testApiPath,
				header:   testHeader,
				query:    testQuery,
				body:     testBody,
			},
			httpHandlerFunc: http.HandlerFunc(
				// configure http stub response
				func(w http.ResponseWriter, r *http.Request) {
					responseInternalServerError := []func(w http.ResponseWriter, r *http.Request){
						// 1st response
						func(w http.ResponseWriter, r *http.Request) {
							http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
						},
						// 2nd response
						func(w http.ResponseWriter, r *http.Request) {
							http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
						},
						// over 3 tries
						func(w http.ResponseWriter, r *http.Request) {
							t.Errorf("Unit Test Configuration Error: Too many retry called.")
						},
					}
					responseInternalServerError[responseCounter](w, r)
					responseCounter++
				}),
			wantRespBody:   "",
			wantStatusCode: 0, // have to modify http.go
			wantErr:        `Get "http://127.0.0.1:8080/apipath?id=hoge": GET http://127.0.0.1:8080/apipath?id=hoge giving up after 2 attempt(s)`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// initialize response counter
			responseCounter = 0
			// create and start server
			l, err := net.Listen("tcp", customURL)
			if err != nil {
				log.Fatalln("Unit Test Configuration Error: Failed to create HTTP stub server:", err)
			}
			ts := httptest.NewUnstartedServer(tt.httpHandlerFunc)
			ts.Listener.Close()
			ts.Listener = l
			ts.Start()
			// exeucte method
			gotRespBody, err, gotStatusCode := tt.h.Request(ts.URL, tt.args.method, tt.args.apipath, tt.args.header, tt.args.query, tt.args.body)
			// close http stub server
			ts.Close()
			// validation
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
