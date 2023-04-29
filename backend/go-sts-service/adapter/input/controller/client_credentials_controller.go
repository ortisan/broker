package controller

import (
	"errors"
	"net/http"
	errApp "ortisan-broker/go-commons/error"
	httpErr "ortisan-broker/go-commons/infrastructure/http/error"
	"ortisan-broker/go-sts-service/adapter/dto"
	"ortisan-broker/go-sts-service/application"

	"github.com/gin-gonic/gin"
)

type CreateClientCredentialsController interface {
	CreateClientCredentials(c *gin.Context)
}

type createClientCredentialsController struct {
	CreateClientCredencialsApplication application.CreateClientCredentialsApplication
}

func NewCreateClientCredentialsController(clientCredentialsApplication application.CreateClientCredentialsApplication) (CreateClientCredentialsController, error) {
	if clientCredentialsApplication == nil {
		return nil, errors.New("clientCredentialsApplication is required")
	}
	return &createClientCredentialsController{
		CreateClientCredencialsApplication: clientCredentialsApplication,
	}, nil
}

func (cccc createClientCredentialsController) CreateClientCredentials(c *gin.Context) {
	var req dto.ClientCredentialsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httpErr.HandleError(c, errApp.NewUnprocessableEntityErrorWithCause("Error to parse body.", err))
		return
	}
	resp, err := cccc.CreateClientCredencialsApplication.CreateClientCredentials(req)
	if err != nil {
		httpErr.HandleError(c, err)
	} else {
		c.JSON(http.StatusCreated, &resp)
	}
}
