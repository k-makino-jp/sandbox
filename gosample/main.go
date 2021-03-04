package main

import (
	"fmt"
	"gosample/pkg/http"
	"time"
)

func httpRequest() {
	// http
	h := http.NewHttpClient(
		2,
		1*time.Second,
		10*time.Second,
		5*time.Second,
	)
	_, err, statusCode := h.Get("https://www.google.com", nil, nil)
	fmt.Println(err, statusCode)
}

func main() {
}
