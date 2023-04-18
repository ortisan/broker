//go:build wireinject
// +build wireinject

package di

import (
	"user-service/adapter/output/database"
	"user-service/application"
	"user-service/domain/usecase"
	"user-service/infrastructure/http/router"

	"github.com/gin-gonic/gin"

	"github.com/google/wire"
)

var RepositoriesSet = wire.NewSet(database.NewCreateUserPostgresRepository, database.NewGetUserPostgresRepository)
var UseCasesSet = wire.NewSet(usecase.NewCreateUserUseCase, usecase.NewGetUserUseCase)

var ApplicationsSet = wire.NewSet(application.NewCreateUserApplication, application.NewGetUserApplication)
var RoutersSet = wire.NewSet(router.NewRouter)
var AppSet = wire.NewSet(RepositoriesSet, UseCasesSet, ApplicationsSet, RoutersSet)

func InitializeRouters() (*gin.Engine, error) {
	wire.Build(AppSet)
	return nil, nil
}
