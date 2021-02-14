// interactor implements usecase
package usecase

import (
	"sandbox/goutil/pkg/01_interface/presenter"
	"sandbox/goutil/pkg/01_interface/repository"
)

func NewSubcmd() *subcmd {
	return &subcmd{
		config: repository.Config{},
		logger: presenter.Logger{},
	}
}

type subcmd struct {
	config repository.ConfigRepository
	logger presenter.LogPresenter
}

func (s subcmd) Setup() {
	s.config.Read()
	s.logger.Errorf("usecase setup")
}

func (s subcmd) Unsetup() {
	s.config.Read()
	s.logger.Errorf("usecase setup")
}

// type ConfigInteractor struct {
// 	configUsecase ConfigUsecase
// }

// func (c *ConfigInteractor) Read(cr ConfigRepository) *config.Config {
// 	config := cr.Read()
// 	return config
// }
