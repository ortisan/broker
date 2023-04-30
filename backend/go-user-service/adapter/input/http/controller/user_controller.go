package controller

import (
	"errors"
	"net/http"
	errApp "ortisan-broker/go-commons/error"
	httpErr "ortisan-broker/go-commons/infrastructure/http/error"
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
	router.POST("/api/v1/users", createRouter.CreateUser)
	return router, nil

}

func (cur *createUserRouter) CreateUser(c *gin.Context) {
	var req dto.User
	if err := c.ShouldBindJSON(&req); err != nil {
		httpErr.HandleError(c, errApp.NewUnprocessableEntityErrorWithCause("Error to parse body.", err))
		return
	}
	resp, err := cur.createUserApplication.CreateUser(c.Request.Context(), req)
	if err != nil {
		httpErr.HandleError(c, err)
	} else {
		c.JSON(http.StatusCreated, &resp)
	}
}

type getByIdParams struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type getUserByIdRouter struct {
	getUserApplication application.GetUserApplication
}

func NewGetUserByIdRouter(router *gin.Engine, getUserApplication application.GetUserApplication) (*gin.Engine, error) {
	if getUserApplication == nil {
		return nil, errors.New("get user application is required")
	}
	getRouter := &getUserByIdRouter{
		getUserApplication: getUserApplication,
	}
	router.GET("/api/v1/users/:id", getRouter.GetUserById)
	return router, nil
}

func (gur *getUserByIdRouter) GetUserById(c *gin.Context) {
	var params getByIdParams
	if err := c.ShouldBindUri(&params); err != nil {
		httpErr.HandleError(c, errApp.NewUnprocessableEntityErrorWithCause("Error to parse params.", err))
		return
	}
	resp, err := gur.getUserApplication.GetUser(c.Request.Context(), params.ID)
	if err != nil {
		httpErr.HandleError(c, err)
	} else {
		c.JSON(http.StatusCreated, &resp)
	}
}
