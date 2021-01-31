package example

import (
	"local.packages/filereader"
	"local.packages/printer"
)

const (
	filepath          = "test.txt"
	exitStatusSuccess = 0
	exitStatusFailed  = 1
)

func execPrinter() {
	p := printer.NewPrinterImpl()
	p.Print()
}

func execFileReader(f filereader.FileReader) uint8 {
	err := f.Read(filepath)
	if err != nil {
		return exitStatusFailed
	}
	return exitStatusSuccess
}

func call() {
	execPrinter()
	f := filereader.NewFileReaderInstance()
	execFileReader(f)
}
