package controller

import (
	"ortisan-broker/go-sts-service/application"

	"github.com/gin-gonic/gin"
)

type TokenController interface {
	CreateToken(c *gin.Context)
}

type tokenController struct {
	createTokenApplication application.CreateClientCredentialsApplication
}
