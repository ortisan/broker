package controller

import (
	"errors"

	"ortisan-broker/go-commons/infrastructure/http/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewClientCredentialsRouter(router *gin.Engine, clientCredentialsController CreateClientCredentialsController) (*gin.Engine, error) {
	if router == nil {
		return nil, errors.New("router is required")
	}
	if clientCredentialsController == nil {
		return nil, errors.New("client credentials controller is required")
	}
	router.POST("/api/v1/clientCredentials", clientCredentialsController.CreateClientCredentials)
	return router, nil
}

func NewRouter(clientCredentialsController CreateClientCredentialsController) (*gin.Engine, error) {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(middleware.RecoveryMiddleware())

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	//r.POST("/client_credentials", GenerateClientCredentials)
	//r.POST("/token", GenerateToken)
	//r.POST("/validate_token", ValidateToken)

	_, err := NewClientCredentialsRouter(router, clientCredentialsController)
	if err != nil {
		return nil, err
	}
	return router, nil
}
