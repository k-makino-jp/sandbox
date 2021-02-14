// controller
// Layer: Interface Adapter: Controllers
// Role: Convert Input (for usecase)
package controller

import (
	usecase "sandbox/goutil/pkg/02_usecase/config"
)

type controller struct{}

func (c controller) subcmd(subcmd string, subcmdUseCase usecase.Subcmd) {
	switch subcmd {
	case "setup":
		subcmdUseCase.Setup()
	case "unsetup":
		subcmdUseCase.Unsetup()
	}
}
