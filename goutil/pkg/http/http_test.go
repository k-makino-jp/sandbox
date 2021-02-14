// http is sample package

package http

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

func Test_httpClient_Get(t *testing.T) {
	customURL := "127.0.0.1:8080"
	testQuery := map[string]string{"id": "hoge"}
	testHeader := map[string]string{"Api-Key": "sample api key"}
	testHttpClient := httpClient{
		retryCount:     5,
		retryWaitMin:   1 * time.Second,
		retryWaitMax:   10 * time.Second,
		requestTimeout: 5 * time.Second,
		client:         *resty.New(),
	}
	var responseCounter uint8
	type args struct {
		baseURL string
		header  map[string]string
		query   map[string]string
	}
	tests := []struct {
		name            string
		h               httpClient
		args            args
		wantRespBody    string
		wantStatusCode  int
		wantErr         error
		httpHandlerFunc http.HandlerFunc
	}{
		{
			name:            "httpClient.Get GetRequest ReturnsErrorEqualsNil",
			h:               testHttpClient,
			args:            args{baseURL: customURL, header: testHeader, query: testQuery},
			httpHandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Hello HTTP Test") }),
			wantRespBody:    "Hello HTTP Test",
			wantStatusCode:  http.StatusOK,
			wantErr:         nil,
		},
		{
			name: "httpClient.Get 404NotFound ReturnsErrorEqualsError",
			h:    testHttpClient,
			args: args{baseURL: customURL, header: testHeader, query: testQuery},
			httpHandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				responseNotFoundError := []func(w http.ResponseWriter, r *http.Request){
					func(w http.ResponseWriter, r *http.Request) { http.Error(w, "404 Not Found.", http.StatusNotFound) },
					func(w http.ResponseWriter, r *http.Request) {
						log.Fatalln("Unit Test Configuration Error: too many retry called.")
					},
				}
				responseNotFoundError[responseCounter](w, r)
				responseCounter++
			}),
			wantRespBody:   "404 Not Found.\n",
			wantStatusCode: http.StatusNotFound,
			wantErr:        nil,
		},
		{
			name: "httpClient.Get 500InternalServerError ReturnsErrorEqualsError",
			h: httpClient{
				retryCount:     1,
				retryWaitMin:   1 * time.Second,
				retryWaitMax:   10 * time.Second,
				requestTimeout: 5 * time.Second,
				client:         *resty.New(),
			},
			args: args{baseURL: customURL, header: testHeader, query: testQuery},
			httpHandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				responseInternalServerError := []func(w http.ResponseWriter, r *http.Request){
					func(w http.ResponseWriter, r *http.Request) {
						http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
					},
					func(w http.ResponseWriter, r *http.Request) {
						http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
					},
				}
				responseInternalServerError[responseCounter](w, r)
				responseCounter++
			}),
			wantRespBody:   "500 Internal Server Error.\n",
			wantStatusCode: http.StatusInternalServerError,
			wantErr:        nil,
		},
		{
			name: "httpClient.Get 2試行目で成功するとき ReturnsErrorEqualsError",
			h: httpClient{
				retryCount:     1,
				retryWaitMin:   1 * time.Second,
				retryWaitMax:   10 * time.Second,
				requestTimeout: 5 * time.Second,
				client:         *resty.New(),
			},
			args: args{baseURL: customURL, header: testHeader, query: testQuery},
			httpHandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				responseInternalServerError := []func(w http.ResponseWriter, r *http.Request){
					func(w http.ResponseWriter, r *http.Request) {
						http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
					},
					func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Hello HTTP Test") },
				}
				responseInternalServerError[responseCounter](w, r)
				responseCounter++
			}),
			wantRespBody:   "Hello HTTP Test",
			wantStatusCode: http.StatusOK,
			wantErr:        nil,
		},
		{
			name:            "httpClient.Get RequestError ReturnsErrorEqualsError",
			h:               testHttpClient,
			args:            args{baseURL: "", header: nil, query: nil},
			httpHandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			wantRespBody:    "",
			wantStatusCode:  0,
			wantErr:         errors.New(`Get "": unsupported protocol scheme ""`),
		},
		{
			name: "httpClient.Get TimeoutError ReturnsErrorEqualsError",
			h: httpClient{
				retryCount:     1,
				retryWaitMin:   1 * time.Second,
				retryWaitMax:   10 * time.Second,
				requestTimeout: 1 * time.Nanosecond,
				client:         *resty.New(),
			},
			args:            args{baseURL: customURL, header: nil, query: nil},
			httpHandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { time.Sleep(2 * time.Nanosecond) }),
			wantRespBody:    "",
			wantStatusCode:  0,
			wantErr:         errors.New(`Get "http://127.0.0.1:8080": context deadline exceeded (Client.Timeout exceeded while awaiting headers)`),
		},
	}
	isSameError := func(err, want error) bool {
		var errString, wantString string
		if err != nil {
			errString = err.Error()
		}
		if want != nil {
			wantString = want.Error()
		}
		if errString == wantString {
			return true
		}
		return false
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseCounter = 0
			// ==================================================
			// create and start server
			// ==================================================
			l, err := net.Listen("tcp", customURL)
			if err != nil {
				log.Fatalln("Unit Test Configuration Error: HTTP Mock Server:", err)
			}
			ts := httptest.NewUnstartedServer(tt.httpHandlerFunc)
			ts.Listener.Close()
			ts.Listener = l
			ts.Start()
			// ==================================================
			// run method
			// ==================================================
			if tt.args.baseURL == customURL {
				tt.args.baseURL = ts.URL
			}
			gotRespBody, err, gotStatusCode := tt.h.Get(tt.args.baseURL, tt.args.header, tt.args.query)
			// ==================================================
			// close server
			// ==================================================
			ts.Close()
			// ==================================================
			// validation
			// ==================================================
			if !isSameError(err, tt.wantErr) {
				t.Errorf("httpClient.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRespBody != tt.wantRespBody {
				t.Errorf("httpClient.Get() gotRespBody = %v, want %v", gotRespBody, tt.wantRespBody)
			}
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("httpClient.Get() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
		})
	}
}

func Test_httpClient_Get_for_Timeout(t *testing.T) {
	customURL := "127.0.0.1:8080"
	testQuery := map[string]string{"id": "hoge"}
	testHeader := map[string]string{"Api-Key": "sample api key"}
	type args struct {
		baseURL string
		header  map[string]string
		query   map[string]string
	}
	tests := []struct {
		name            string
		h               httpClient
		args            args
		wantRespBody    string
		wantStatusCode  int
		wantErr         error
		wantElapsed     time.Duration
		httpHandlerFunc http.HandlerFunc
	}{
		{
			name: "httpClient.Get RetryTimeValidation ReturnsErrorEqualsError",
			h: httpClient{
				retryCount:     3, // 1, 2, 4, 8, 16 → 2, 4, 4, ...
				retryWaitMin:   2 * time.Second,
				retryWaitMax:   4 * time.Second,
				requestTimeout: 5 * time.Second,
				client:         *resty.New(),
			},
			args: args{baseURL: customURL, header: testHeader, query: testQuery},
			httpHandlerFunc: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			}),
			wantRespBody:   "500 Internal Server Error.\n",
			wantStatusCode: http.StatusInternalServerError,
			wantErr:        nil,
			wantElapsed:    10 * time.Second,
		},
	}
	isSameError := func(err, want error) bool {
		var errString, wantString string
		if err != nil {
			errString = err.Error()
		}
		if want != nil {
			wantString = want.Error()
		}
		if errString == wantString {
			return true
		}
		return false
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ==================================================
			// create and start server
			// ==================================================
			l, err := net.Listen("tcp", customURL)
			if err != nil {
				log.Fatalln("Unit Test Configuration Error: HTTP Mock Server:", err)
			}
			ts := httptest.NewUnstartedServer(tt.httpHandlerFunc)
			ts.Listener.Close()
			ts.Listener = l
			ts.Start()
			// ==================================================
			// run method
			// ==================================================
			fmt.Printf("Please wait about %v...\n", tt.wantElapsed)
			start := time.Now()
			gotRespBody, err, gotStatusCode := tt.h.Get(ts.URL, tt.args.header, tt.args.query)
			gotElapsed := time.Since(start).Truncate(1 * time.Second) // ミリ秒以下切り捨て
			// ==================================================
			// close server
			// ==================================================
			ts.Close()
			// ==================================================
			// validation
			// ==================================================
			if gotElapsed.Seconds() != tt.wantElapsed.Seconds() {
				t.Errorf("httpClient.Get() gotElapsed = %v, wantElapsed %v", gotElapsed, tt.wantElapsed)
			}
			if !isSameError(err, tt.wantErr) {
				t.Errorf("httpClient.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRespBody != tt.wantRespBody {
				t.Errorf("httpClient.Get() gotRespBody = %v, want %v", gotRespBody, tt.wantRespBody)
			}
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("httpClient.Get() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
		})
	}
}

func TestNewHttpClient(t *testing.T) {
	type args struct {
		retryCount     int
		retryWaitMin   time.Duration
		retryWaitMax   time.Duration
		requestTimeout time.Duration
	}
	tests := []struct {
		name string
		args args
		want *httpClient
	}{
		{
			name: "NewHttpClient インスタンスを生成するとき インスタンスを返すこと",
			args: args{
				retryCount:     5,
				retryWaitMin:   1 * time.Second,
				retryWaitMax:   10 * time.Second,
				requestTimeout: 5 * time.Second,
			},
			want: &httpClient{
				retryCount:     5,
				retryWaitMin:   1 * time.Second,
				retryWaitMax:   10 * time.Second,
				requestTimeout: 5 * time.Second,
				client:         resty.Client{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restyNew = func() *resty.Client {
				return &resty.Client{}
			}
			if got := NewHttpClient(tt.args.retryCount, tt.args.retryWaitMin, tt.args.retryWaitMax, tt.args.requestTimeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHttpClient() = %v, want %v", got, tt.want)
			}
			restyNew = resty.New
		})
	}
}
