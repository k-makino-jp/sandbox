// Frameworks and Drivers: cli
// Interface Adapters: usecase caller
// Application Business Rules: usecase
// Enterprise Business Rules: service, model
package main

import (
	"fmt"
	"sandbox/goutil/pkg/01_interface/subcmd"
	"sandbox/goutil/pkg/02_usecase/http"
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
	httpRequest()
	subcmd.HttpRequest()
}
