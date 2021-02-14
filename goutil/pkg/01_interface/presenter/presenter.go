package presenter

import "fmt"

type LogPresenter interface {
	Errorf(string)
}

type Logger struct{}

func (l Logger) Errorf(format string) {
	fmt.Println(format)
}
