// printer is sample package
package printer

import (
	"fmt"
	"gosample/hoge/enver"
)

type Printer interface {
	Print()
}

type printImpl struct {
	enver enver.Enver
}

func (p *printImpl) Print() {
	p.enver.Initialize()
	e := p.enver.GetEnv()
	fmt.Println(e)
}

func NewPrintImpl() *printImpl {
	return &printImpl{
		enver.NewEnvImpl(),
	}
}
