package middleware

import (
	"net/http"
	"ortisan-broker/go-commons/adapter/dto"
	errApp "ortisan-broker/go-commons/error"

	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
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
