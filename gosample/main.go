package main

import (
	"fmt"
	"time"
	// "gosample/hoge/configer"

	"gosample/hoge/http"
	"gosample/hoge/printer"
	// "gosample/hoge/scanner"
)

func sample() {
	// create instance
	p := printer.NewPrintImpl()
	// execute function
	p.Print()

	// create instance
	// configFilePath := "config.json"
	// c := configer.NewConfiger(configFilePath)
	// c.Read()
	// config := c.Get()
	// fmt.Println(config)

	// // timerr
	// t := timer.NewTimerImpl()
	// fmt.Println(t.Now())
	// fmt.Println(t.NowRFC3389())

	// scanner
	// s := scanner.NewScannerImpl("scandata.txt")
	// line, err := s.Scan()
	// if err != nil {
	// 	log.Fatal("ERROR", line)
	// }
	// fmt.Println(s.Get())

	// http
	// retryMax := 3
	// retryWaitMin := 1 * time.Second
	// retryWaitMax := 5 * time.Second
	// httpRequestTimeout := 5 * time.Second
	// h := http.NewHttpClientImpl(retryMax, retryWaitMin, retryWaitMax, httpRequestTimeout)
	// endpoint := "https://www.google.com/"
	// apipath := ""
	// var header, query map[string]string
	// respBody, err, statusCode := h.Request(endpoint, "GET", apipath, header, query, nil)
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println("=== StatusCode =", statusCode)
	// fmt.Println(respBody)
	// _, err, statusCode := h.Get(endpoint, header, query)
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println("=== StatusCode =", statusCode)
	// fmt.Println(string(respBody))
}

func main() {
	// encryptor
	// input := "plaintext.json"
	// output := "encrypted.json"
	// e := encrypt.NewEncryptorImpl(input, output)
	// e.Encrypt()

	// // decryptor
	// input = "encrypted.json"
	// output = "decrypted.json"
	// d := decrypt.NewDecryptorImpl(input, output)
	// // The key should be 32 bytes (AES-256)
	// key := []byte("12345678901234567890123456789012")
	// d.Decrypt(key)

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
