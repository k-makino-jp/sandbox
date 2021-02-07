package main

import (
	"fmt"
	"gosample/hoge/configer"
	"gosample/hoge/decryptor"
	"gosample/hoge/encryptor"
	"gosample/hoge/http"
	"gosample/hoge/printer"
	"log"
	"time"

	// "gosample/hoge/scanner"
	timer "gosample/hoge/timer"
)

func sample() {
	// create instance
	p := printer.NewPrintImpl()
	// execute function
	p.Print()

	// create instance
	configFilePath := "config.json"
	c := configer.NewConfigerImpl(configFilePath)
	c.Read()
	config := c.Get()
	fmt.Println(config)

	// timerr
	t := timer.NewTimerImpl()
	fmt.Println(t.Now())
	fmt.Println(t.NowRFC3389())

	// scanner
	// s := scanner.NewScannerImpl("scandata.txt")
	// line, err := s.Scan()
	// if err != nil {
	// 	log.Fatal("ERROR", line)
	// }
	// fmt.Println(s.Get())

	// http
	retryMax := 3
	retryWaitMin := 1 * time.Second
	retryWaitMax := 5 * time.Second
	httpRequestTimeout := 5 * time.Second
	h := http.NewHttpClientImpl(retryMax, retryWaitMin, retryWaitMax, httpRequestTimeout)
	endpoint := "https://www.google.com/"
	apipath := ""
	var header, query map[string]string
	respBody, err, statusCode := h.Request(endpoint, "GET", apipath, header, query, nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("=== StatusCode =", statusCode)
	fmt.Println(respBody)
	// _, err, statusCode := h.Get(endpoint, header, query)
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println("=== StatusCode =", statusCode)
	// fmt.Println(string(respBody))
}

func main() {
	// encryptor
	input := "plaintext.json"
	output := "encrypted.json"
	e := encryptor.NewEncryptorImpl(input, output)
	e.Encrypt()

	// decryptor
	input = "encrypted.json"
	output = "decrypted.json"
	d := decryptor.NewDecryptorImpl(input, output)
	// The key should be 32 bytes (AES-256)
	key := []byte("12345678901234567890123456789012")
	d.Decrypt(key)
}
