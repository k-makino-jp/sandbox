package main

import (
	"gosample/hoge/printer"
)

func main() {
	// create instance
	p := printer.NewPrintImpl()
	// execute function
	p.Print()
}
