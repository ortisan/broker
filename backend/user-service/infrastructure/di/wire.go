//go:build wireinject
// +build wireinject

package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"user-service/adapter/output/database"
	"user-service/application"
	"user-service/config"
	"user-service/domain/usecase"
	"user-service/infrastructure/datastore"
	"user-service/infrastructure/http/router"
	"user-service/infrastructure/log"
)

var ConfigSet = wire.NewSet(config.NewConfig)
var LoggerSet = wire.NewSet(log.NewLogger)
var DbSet = wire.NewSet(datastore.NewDB)
var RepositoriesSet = wire.NewSet(database.NewCreateUserPostgresRepository, database.NewGetUserPostgresRepository)
var UseCasesSet = wire.NewSet(usecase.NewCreateUserUseCase, usecase.NewGetUserUseCase)
var ApplicationsSet = wire.NewSet(application.NewCreateUserApplication, application.NewGetUserApplication)
var RoutersSet = wire.NewSet(router.NewRouter)
var AppSet = wire.NewSet(ConfigSet, DbSet, RepositoriesSet, UseCasesSet, ApplicationsSet, RoutersSet)

func ConfigRouters() (*gin.Engine, error) {
	wire.Build(AppSet)
	return nil, nil
}
