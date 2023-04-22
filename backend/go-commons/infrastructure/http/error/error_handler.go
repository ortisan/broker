package error

import (
	"net/http"
	"ortisan-broker/go-commons/adapter/dto"
	errApp "ortisan-broker/go-commons/error"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	switch err.(type) {
	case *errApp.BadArgumentError:
		error := err.(*errApp.BadArgumentError)
		c.JSON(http.StatusBadRequest, dto.Error{Message: error.Error(), Cause: error.Cause(), StackTrace: error.StackTrace()})
	case *errApp.AuthError:
		error := err.(*errApp.AuthError)
		c.JSON(http.StatusUnauthorized, dto.Error{Message: error.Error(), Cause: error.Cause(), StackTrace: error.StackTrace()})
	case *errApp.NotFoundError:
		error := err.(*errApp.NotFoundError)
		c.JSON(http.StatusNotFound, dto.Error{Message: error.Error(), Cause: error.Cause(), StackTrace: error.StackTrace()})
	case *errApp.ConflictError:
		error := err.(*errApp.ConflictError)
		c.JSON(http.StatusConflict, dto.Error{Message: error.Error(), Cause: error.Cause(), StackTrace: error.StackTrace()})
	case *errApp.UnprocessableEntityError:
		error := err.(*errApp.UnprocessableEntityError)
		c.JSON(http.StatusUnprocessableEntity, dto.Error{Message: error.Error(), Cause: error.Cause(), StackTrace: error.StackTrace()})
	default:
		c.JSON(http.StatusInternalServerError, dto.Error{Message: err.(error).Error(), StackTrace: string(debug.Stack())})
	}
}
