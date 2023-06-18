package http

import (
	"errors"
	"net/http"
	errapp "ortisan-broker/go-commons/error"
	httperr "ortisan-broker/go-commons/infrastructure/http/error"
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
		httperr.HandleError(c, errapp.NewUnprocessableEntityErrorWithCause("Error to parse body.", err))
		return
	}
	resp, err := cccc.CreateClientCredencialsApplication.CreateClientCredentials(c.Request.Context(), &req)
	if err != nil {
		httperr.HandleError(c, err)
	} else {
		c.JSON(http.StatusCreated, &resp)
	}
}