package main

import (
	"fmt"
	"gosample/pkg/http"
	"runtime"
	"time"
)

func stackTrace(a int) bool {
	pc, _, line, ok := runtime.Caller(0)
	if !ok {
		return ok
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Printf("line=%d, func=%v\n", line, funcName)
	stackTrace1()
	return false
}

func stackTrace1() bool {
	pc, _, line, ok := runtime.Caller(1)
	if !ok {
		return ok
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Printf("line=%d, func=%v\n", line, funcName)
	return false
}

func stackTrace2() bool {
	pc, _, line, ok := runtime.Caller(2)
	if !ok {
		return ok
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Printf("line=%d, func=%v\n", line, funcName)
	return false
}

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
	stackTrace(1)
	stackTrace1()
	stackTrace2()
}
