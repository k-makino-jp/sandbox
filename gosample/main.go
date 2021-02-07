package main

import (
	"fmt"
	"gosample/hoge/configer"
	"gosample/hoge/decryptor"
	"gosample/hoge/encryptor"
	"gosample/hoge/http"
	"gosample/hoge/printer"
	"log"

	// "gosample/hoge/scanner"
	time "gosample/hoge/timer"
)

func main() {
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

	// encryptor
	outputFilePath := "encypted.json"
	e := encryptor.NewEncryptorImpl(outputFilePath)
	e.Encrypt()

	// decryptor
	outputFilePath = "decrypted.json"
	d := decryptor.NewDecryptorImpl("encypted.json", outputFilePath)
	d.Decrypt()

	// scanner
	// s := scanner.NewScannerImpl("scandata.txt")
	// line, err := s.Scan()
	// if err != nil {
	// 	log.Fatal("ERROR", line)
	// }
	// fmt.Println(s.Get())

	// timer
	t := time.NewTimerImpl()
	fmt.Println(t.Now())
	fmt.Println(t.NowRFC3389())

	// http
	h := http.NewHttpClientImpl()
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
