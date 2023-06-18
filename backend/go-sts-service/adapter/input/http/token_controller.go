package http

import (
	"errors"
	"net/http"
	errapp "ortisan-broker/go-commons/error"
	httpErr "ortisan-broker/go-commons/infrastructure/http/error"
	"ortisan-broker/go-sts-service/adapter/dto"
	"ortisan-broker/go-sts-service/application"

	"github.com/gin-gonic/gin"
)

type OauthTokenController interface {
	CreateOauthToken(c *gin.Context)
}

type oauthTokenController struct {
	createTokenApplication application.CreateOauthTokenApplication
}

func (t *oauthTokenController) CreateOauthToken(c *gin.Context) {
	var req dto.OauthTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httpErr.HandleError(c, errapp.NewUnprocessableEntityErrorWithCause("Error to parse body.", err))
		return
	}
	resp, err := t.createTokenApplication.CreateOauthToken(c.Request.Context(), &req)
	if err != nil {
		httpErr.HandleError(c, err)
	} else {
		c.JSON(http.StatusCreated, &resp)
	}
}

func NewOauthTokenController(createTokenApplication application.CreateOauthTokenApplication) (OauthTokenController, error) {
	if createTokenApplication == nil {
		return nil, errors.New("create token application is required")
	}

	return &oauthTokenController{
		createTokenApplication: createTokenApplication,
	}, nil
}
