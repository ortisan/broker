//go:build wireinject
// +build wireinject

package di

import (
	"ortisan-broker/go-commons/config"
	"ortisan-broker/go-commons/infrastructure/log"
	"ortisan-broker/go-sts-service/adapter/input/controller"
	"ortisan-broker/go-sts-service/application"
	"ortisan-broker/go-sts-service/domain/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ConfigSet = wire.NewSet(config.NewConfig)
var LoggerSet = wire.NewSet(log.NewLogger)
var UseCasesSet = wire.NewSet(usecase.NewCreateClientCredentialsUseCase, usecase.NewGetClientCredentialsUseCase, usecase.NewCreateOauthTokenUseCase)
var ApplicationsSet = wire.NewSet(application.NewClientCredentialsAdapter, application.NewCreateClientCredentialsApplication, application.NewOauthTokenAdapter, application.NewCreateOauthTokenApplication)
var ControllersSet = wire.NewSet(controller.NewCreateClientCredentialsController, controller.NewOauthTokenController)
var RoutersSet = wire.NewSet(controller.NewRouter)
var AppSet = wire.NewSet(ConfigSet, LoggerSet, UseCasesSet, ApplicationsSet, ControllersSet, RoutersSet)

func ConfigRouters() (*gin.Engine, error) {
	wire.Build(AppSet)
	return nil, nil
}
