package http

import (
	"errors"
	"go.opentelemetry.io/otel"
	"net/http"
	errapp "ortisan-broker/go-commons/error"
	httperr "ortisan-broker/go-commons/infrastructure/http/error"
	"ortisan-broker/go-sts-service/adapter/dto"
	"ortisan-broker/go-sts-service/application"
	"reflect"
	"runtime"

	"github.com/gin-gonic/gin"
)

type OauthTokenController interface {
	CreateOauthToken(c *gin.Context)
}

type oauthTokenController struct {
	createTokenApplication application.CreateOauthTokenApplication
}

func (otc *oauthTokenController) CreateOauthToken(c *gin.Context) {
	pc, _, _, _ := runtime.Caller(0)
	newCtx, span := otel.Tracer(reflect.TypeOf(otc).String()).Start(c, runtime.FuncForPC(pc).Name())
	defer span.End()

	var req dto.OauthTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httperr.HandleError(c, errapp.NewUnprocessableEntityErrorWithCause("Error to parse body.", err))
		return
	}
	resp, err := otc.createTokenApplication.CreateOauthToken(newCtx, &req)
	if err != nil {
		httperr.HandleError(c, err)
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
