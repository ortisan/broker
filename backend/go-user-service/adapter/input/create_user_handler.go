package handler

import (
	"net/http"
	errorApp "ortisan-broker/go-commons/error"
	"ortisan-broker/go-user-service/adapter/dto"
	"ortisan-broker/go-user-service/application/user"

	"github.com/gin-gonic/gin"
)

type CreateUserGinHandler interface {
	CreateUser(c *gin.Context)
}

type createUserHandler struct {
	createUserApplication user.CreateUserApplication
}

func (cua createUserHandler) CreateUser(c *gin.Context) {
	var user dto.User
	if err := c.ShouldBindJSON(&user); err != nil {
		panic(errorApp.NewBadArgumentErrorWithCause("Error to parse body.", err))
	}
	resp, err := cua.createUserApplication.CreateUser(user)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusCreated, &resp)
	}
}
