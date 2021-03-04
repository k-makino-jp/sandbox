package main

import (
	"gosample/pkg/azure"
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
	message := azure.Message{
		Data: "data",
	}
	a := azure.NewAzure()
	a.InitPipeline(azure.Sas{})
	a.Enqueue(message)
}

// func KubernetesAPI() {
// 	k, err := kubernetes.NewKubernetes()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	info, err := k.GetInfo()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(info)
// }

func main() {
	AzureEnqueue()
	// KubernetesAPI()
}
