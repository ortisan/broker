package router

import (
	"net/http"
	errApp "ortisan-broker/go-commons/error"
	"ortisan-broker/go-user-service/adapter/dto"
	"ortisan-broker/go-user-service/application"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case errApp.BadArgumentError:
					error := err.(errApp.BadArgumentError)
					c.JSON(http.StatusBadRequest, dto.Error{Message: error.Error(), Cause: error.Cause(), StackTrace: error.StackTrace()})
				case errApp.AuthError:
					error := err.(errApp.AuthError)
					c.JSON(http.StatusUnauthorized, dto.Error{Message: error.Error(), Cause: error.Cause(), StackTrace: error.StackTrace()})
				case errApp.NotFoundError:
					error := err.(errApp.NotFoundError)
					c.JSON(http.StatusNotFound, dto.Error{Message: error.Error(), Cause: error.Cause(), StackTrace: error.StackTrace()})
				default:
					c.JSON(http.StatusInternalServerError, dto.Error{Message: err.(error).Error(), StackTrace: string(debug.Stack())})
				}
			}
		}()
		c.Next()
	}
}

func NewRouter(createUserApplication application.CreateUserApplication, getUserApplication application.GetUserApplication) (*gin.Engine, error) {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(recovery())
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
