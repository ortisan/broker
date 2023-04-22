package controller

import (
	"net/http"

	"ortisan-broker/go-sts-service/dto"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

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
