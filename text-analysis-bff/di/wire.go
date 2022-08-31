//go:build wireinject
// +build wireinject

package di

import (
	"echo-get-started/driver"
	"echo-get-started/rest/handler"
	"echo-get-started/use_case"

	"github.com/google/wire"
)

func InitNounHandler() *handler.NounHandler {
	wire.Build(handler.NewNounHandler, use_case.NewNounUseCase, driver.NewNounPort)
	return nil
}
