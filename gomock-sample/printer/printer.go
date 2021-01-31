package printer

import (
	"fmt"
)

type Printer interface {
	Print()
}

type printerImpl struct {
}

// pointer receiver can handle own variables
func (p *printerImpl) Print() {
	fmt.Println("Call: Print()")
}

func NewPrinterImpl() *printerImpl {
	return &printerImpl{}
}
