package main

import (
	"fmt"
	"integtest/pkg/http"
)

func main() {
	get()
	put()
	post()
}

func get() {
	queries := map[string]string{
		"api_key":    "api-key-here",
		"api_secert": "api-secert",
	}
	headers := map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   "My custom User Agent String",
	}
	url := "https://httpbin.org/get"
	resp, err := http.Get(queries, headers, url)

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
}

func put() {
	queries := map[string]string{
		"api_key":    "api-key-here",
		"api_secert": "api-secert",
	}
	headers := map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   "My custom User Agent String",
	}
	body := `{"username":"testuser", "password":"testpass"}`
	url := "https://httpbin.org/put"
	resp, err := http.Put(queries, headers, body, url)

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
}

func post() {
	queries := map[string]string{
		"api_key":    "api-key-here",
		"api_secert": "api-secert",
	}
	headers := map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   "My custom User Agent String",
	}
	body := `{"username":"testuser", "password":"testpass"}`
	url := "https://httpbin.org/post"
	resp, err := http.Post(queries, headers, body, url)

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
}
