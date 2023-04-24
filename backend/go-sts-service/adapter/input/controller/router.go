package controller

import (
	"net/http"

	"ortisan-broker/go-commons/infrastructure/http/middleware"
	"ortisan-broker/go-sts-service/dto"
	"ortisan-broker/go-user-service/application"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

func GenerateClientCredentials(c *gin.Context) {
	var req dto.ClientCredentialsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(errSts.NewBadRequestError("Error to parse body.", err))
	}
	cc, err := gss.GenerateClientCredentials(req)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, cc)
	}
}

func GenerateToken(c *gin.Context) {
	var req dto.TokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(errSts.NewBadRequestError("Error to parse body.", err))
	}
	token, err := gss.GenerateToken(req)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, dto.TokenResponse{Token: token})
	}
}

func ValidateToken(c *gin.Context) {
	var req dto.ValidateTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(errSts.NewBadRequestError("Error to parse body.", err))
	}
	resp, err := gss.ValidateToken(req)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func CreateRoutes(ss service.StsService) {

	gss = ss

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(Recovery())

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.POST("/client_credentials", GenerateClientCredentials)
	r.POST("/token", GenerateToken)
	r.POST("/validate_token", ValidateToken)

	r.Run(":8080")
}
