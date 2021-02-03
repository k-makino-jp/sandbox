package main

import (
	"fmt"
	"gosample/hoge/configer"
	"gosample/hoge/decryptor"
	"gosample/hoge/encryptor"
	"gosample/hoge/printer"
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
}
