package router

import (
	"net/http"
	errApp "ortisan-broker/go-commons/error"
	"ortisan-broker/go-user-service/adapter/dto"
	"ortisan-broker/go-user-service/application"

	"github.com/gin-gonic/gin"
)

type createUserRouter struct {
	createUserApplication application.CreateUserApplication
}

func NewCreateUserRouter(router *gin.Engine, createUserApplication application.CreateUserApplication) (*gin.Engine, error) {
	if createUserApplication == nil {
		return nil, errApp.NewBadArgumentError("create user application is required")
	}
	createRouter := &createUserRouter{
		createUserApplication: createUserApplication,
	}
	router.POST("/api/users", createRouter.CreateUser)
	return router, nil

}

func (cur *createUserRouter) CreateUser(c *gin.Context) {
	var req dto.User
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(errApp.NewBadArgumentErrorWithCause("Error to parse body.", err))
	}
	resp, err := cur.createUserApplication.CreateUser(req)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusCreated, &resp)
	}
}