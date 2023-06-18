//go:build wireinject
// +build wireinject

package di

import (
	"ortisan-broker/go-commons/config"
	infraDb "ortisan-broker/go-commons/infrastructure/database"
	"ortisan-broker/go-commons/infrastructure/log"
	"ortisan-broker/go-sts-service/adapter/input/http"
	"ortisan-broker/go-sts-service/adapter/output/database"
	"ortisan-broker/go-sts-service/application"
	"ortisan-broker/go-sts-service/domain/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ConfigSet = wire.NewSet(config.NewConfig)
var LoggerSet = wire.NewSet(log.NewLogger)
var DbSet = wire.NewSet(infraDb.NewDB)
var RepositoriesSet = wire.NewSet(database.NewClientCredentialsAdapter, database.NewClientCredentialsPostgresRepository)
var UseCasesSet = wire.NewSet(usecase.NewCreateClientCredentialsUseCase, usecase.NewGetClientCredentialsUseCase, usecase.NewCreateOauthTokenUseCase)
var ApplicationsSet = wire.NewSet(application.NewClientCredentialsAdapter, application.NewCreateClientCredentialsApplication, application.NewOauthTokenAdapter, application.NewCreateOauthTokenApplication)
var ControllersSet = wire.NewSet(http.NewCreateClientCredentialsController, http.NewOauthTokenController)
var RoutersSet = wire.NewSet(http.NewRouter)
var AppSet = wire.NewSet(ConfigSet, LoggerSet, DbSet, RepositoriesSet, UseCasesSet, ApplicationsSet, ControllersSet, RoutersSet)

func ConfigRouters() (*gin.Engine, error) {
	wire.Build(AppSet)
	return nil, nil
}
