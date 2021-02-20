package main

import (
	"gosample/pkg/azurev2"
)

// func httpRequest() {
// 	// http
// 	h := http.NewHttpClient(
// 		2,
// 		1*time.Second,
// 		10*time.Second,
// 		5*time.Second,
// 	)
// 	_, err, statusCode := h.Get("https://www.google.com", nil, nil)
// 	fmt.Println(err, statusCode)
// }

func AzureEnqueue() {
	message := azurev2.Message{
		Data: "data",
	}
	azure := azurev2.NewAzure(azurev2.Sas{})
	azure.Enqueue(message)
}

func main() {
	AzureEnqueue()
}
