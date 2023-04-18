//go:build wireinject
// +build wireinject

package di

import (
	"monolith/adapter/output/database"
	"monolith/application"
	"monolith/domain/usecase"
	"monolith/infrastructure/http/router"

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
