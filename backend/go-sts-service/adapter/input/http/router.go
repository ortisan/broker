package http

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

func NewOauthTokenRouter(router *gin.Engine, oauthTokenController OauthTokenController) (*gin.Engine, error) {
	if router == nil {
		return nil, errors.New("router is required")
	}
	if oauthTokenController == nil {
		return nil, errors.New("oauth token controller is required")
	}
	router.POST("/api/v1/oauthToken", oauthTokenController.CreateOauthToken)
	return router, nil
}

func NewRouter(clientCredentialsController CreateClientCredentialsController, oauthTokenController OauthTokenController) (*gin.Engine, error) {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(middleware.RecoveryMiddleware())

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	if _, err := NewClientCredentialsRouter(router, clientCredentialsController); err != nil {
		return nil, err
	}
	if _, err := NewOauthTokenRouter(router, oauthTokenController); err != nil {
		return nil, err
	}

	return router, nil
}
