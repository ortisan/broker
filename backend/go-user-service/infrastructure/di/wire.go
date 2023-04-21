//go:build wireinject
// +build wireinject

package di

import (
	database "ortisan-broker/go-user-service/adapter/output/database/user"
	"ortisan-broker/go-user-service/application"
	"ortisan-broker/go-user-service/config"
	"ortisan-broker/go-user-service/domain/usecase"
	datastore "ortisan-broker/go-user-service/infrastructure/database"
	"ortisan-broker/go-user-service/infrastructure/http/router"
	"ortisan-broker/go-user-service/infrastructure/log"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
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
