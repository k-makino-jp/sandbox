// http is sample package

package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_httpClientImpl_Get(t *testing.T) {
	httpMockServer := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello HTTP Test")
		},
	)
	// configure variables
	testQuery := map[string]string{
		"id": "hoge",
	}
	testHeader := map[string]string{
		"Api-Key": "sample api key",
	}
	type args struct {
		baseURL string
		header  map[string]string
		query   map[string]string
	}
	tests := []struct {
		name           string
		h              *httpClientImpl
		args           args
		httpMockServer *httptest.Server
		wantRespBody   []byte
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "httpClinetImpl.Get GetRequest ReturnsErrorEqualsNil",
			h:    &httpClientImpl{},
			args: args{
				baseURL: "",
				header:  testHeader,
				query:   testQuery,
			},
			httpMockServer: httptest.NewServer(
				http.HandlerFunc(
					func(w http.ResponseWriter, r *http.Request) {
						fmt.Fprintf(w, "Hello HTTP Test")
					},
				),
			),
			wantRespBody:   []byte("Hello HTTP Test"),
			wantStatusCode: 200,
			wantErr:        false,
		},
		// {
		// 	name: "httpClinetImpl.Get GetRequest ReturnsErrorEqualsNil",
		// 	h:    &httpClientImpl{},
		// 	args: args{
		// 		baseURL: "",
		// 		header:  testHeader,
		// 		query:   testQuery,
		// 	},
		// 	httpMockServer: httptest.NewServer(
		// 		http.HandlerFunc(
		// 			func(w http.ResponseWriter, r *http.Request) {
		// 				http.Error(w, "Not Found.", http.StatusNotFound)
		// 			},
		// 		),
		// 	),
		// 	wantRespBody:   []byte("Not Found."),
		// 	wantStatusCode: 404,
		// 	wantErr:        false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(httpMockServer)
			defer ts.Close()
			var baseURL string
			if tt.args.baseURL == "" {
				baseURL = ts.URL
			} else {
				baseURL = tt.args.baseURL
			}
			gotRespBody, err, gotStatusCode := tt.h.Get(baseURL, tt.args.header, tt.args.query)
			if (err != nil) != tt.wantErr {
				fmt.Println("ERROR")
				t.Errorf("httpClientImpl.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRespBody, tt.wantRespBody) {
				t.Errorf("httpClientImpl.Get() gotRespBody = %v, want %v", gotRespBody, tt.wantRespBody)
			}
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("httpClientImpl.Get() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
		})
	}
}

func TestNewHttpClientImpl(t *testing.T) {
	tests := []struct {
		name string
		want *httpClientImpl
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHttpClientImpl(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHttpClientImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}
