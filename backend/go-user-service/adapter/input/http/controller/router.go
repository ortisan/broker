package controller

import (
	"ortisan-broker/go-commons/infrastructure/http/middleware"
	"ortisan-broker/go-user-service/application"

	"github.com/gin-gonic/gin"
)

func NewRouter(createUserApplication application.CreateUserApplication, getUserApplication application.GetUserApplication) (*gin.Engine, error) {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(middleware.RecoveryMiddleware())
	_, errorCreate := NewCreateUserRouter(r, createUserApplication)
	if errorCreate != nil {
		return nil, errorCreate
	}
	_, errorGetUser := NewGetUserByIdRouter(r, getUserApplication)
	if errorGetUser != nil {
		return nil, errorGetUser
	}
	return r, nil
}
