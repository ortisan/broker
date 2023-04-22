//go:build wireinject
// +build wireinject

package di

import (
	"ortisan-broker/go-commons/infrastructure/log"
	"ortisan-broker/go-user-service/adapter/input/http/controller"
	outDb "ortisan-broker/go-user-service/adapter/output/database"
	"ortisan-broker/go-user-service/application"
	"ortisan-broker/go-user-service/config"
	"ortisan-broker/go-user-service/domain/usecase"
	infraDb "ortisan-broker/go-user-service/infrastructure/database"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ConfigSet = wire.NewSet(config.NewConfig)
var LoggerSet = wire.NewSet(log.NewLogger)
var DbSet = wire.NewSet(infraDb.NewDB)
var RepositoriesSet = wire.NewSet(outDb.NewCreateUserPostgresRepository, outDb.NewGetUserPostgresRepository)
var UseCasesSet = wire.NewSet(usecase.NewCreateUserUseCase, usecase.NewGetUserUseCase)
var ApplicationsSet = wire.NewSet(application.NewCreateUserApplication, application.NewGetUserApplication)
var RoutersSet = wire.NewSet(controller.NewRouter)
var AppSet = wire.NewSet(ConfigSet, LoggerSet, DbSet, RepositoriesSet, UseCasesSet, ApplicationsSet, RoutersSet)

func ConfigRouters() (*gin.Engine, error) {
	wire.Build(AppSet)
	return nil, nil
}
