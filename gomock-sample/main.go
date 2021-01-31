package main

import (
	"local.packages/filereader"
	"local.packages/printer"
)

func execPrinter() {
	p := printer.NewPrinterImpl()
	p.Print()
}

func execFileReader() {
	f := filereader.NewFileReaderImpl()
	f.Read("test.txt")
}

func main() {
	execPrinter()
	execFileReader()
}
