// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"ortisan-broker/go-commons/infrastructure/log"
	"ortisan-broker/go-user-service/adapter/input/http/controller"
	"ortisan-broker/go-user-service/adapter/output/database"
	"ortisan-broker/go-user-service/application"
	"ortisan-broker/go-user-service/config"
	"ortisan-broker/go-user-service/domain/usecase"
	"ortisan-broker/go-user-service/infrastructure/database"
)

// Injectors from wire.go:

func ConfigRouters() (*gin.Engine, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	db, err := datastore.NewDB(configConfig)
	if err != nil {
		return nil, err
	}
	createUser, err := database.NewCreateUserPostgresRepository(db)
	if err != nil {
		return nil, err
	}
	usecaseCreateUser, err := usecase.NewCreateUserUseCase(createUser)
	if err != nil {
		return nil, err
	}
	createUserApplication, err := application.NewCreateUserApplication(usecaseCreateUser)
	if err != nil {
		return nil, err
	}
	getUser, err := database.NewGetUserPostgresRepository(db)
	if err != nil {
		return nil, err
	}
	usecaseGetUser, err := usecase.NewGetUserUseCase(getUser)
	if err != nil {
		return nil, err
	}
	getUserApplication, err := application.NewGetUserApplication(usecaseGetUser)
	if err != nil {
		return nil, err
	}
	engine, err := controller.NewRouter(createUserApplication, getUserApplication)
	if err != nil {
		return nil, err
	}
	return engine, nil
}

// wire.go:

var ConfigSet = wire.NewSet(config.NewConfig)

var LoggerSet = wire.NewSet(log.NewLogger)

var DbSet = wire.NewSet(datastore.NewDB)

var RepositoriesSet = wire.NewSet(database.NewCreateUserPostgresRepository, database.NewGetUserPostgresRepository)

var UseCasesSet = wire.NewSet(usecase.NewCreateUserUseCase, usecase.NewGetUserUseCase)

var ApplicationsSet = wire.NewSet(application.NewCreateUserApplication, application.NewGetUserApplication)

var RoutersSet = wire.NewSet(controller.NewRouter)

var AppSet = wire.NewSet(ConfigSet, DbSet, RepositoriesSet, UseCasesSet, ApplicationsSet, RoutersSet)