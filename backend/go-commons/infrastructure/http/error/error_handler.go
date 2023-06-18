package error

import (
	"net/http"
	"ortisan-broker/go-commons/adapter/dto"
	errapp "ortisan-broker/go-commons/error"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	switch err.(type) {
	case *errapp.BadArgumentError:
		error := err.(*errapp.BadArgumentError)
		c.JSON(http.StatusBadRequest, dto.Error{Message: error.Error(), Cause: error.Cause(), StackTrace: error.StackTrace()})
	case *errapp.AuthError:
		error := err.(*errapp.AuthError)
		c.JSON(http.StatusUnauthorized, dto.Error{Message: error.Error(), Cause: error.Cause(), StackTrace: error.StackTrace()})
	case *errapp.NotFoundError:
		error := err.(*errapp.NotFoundError)
		c.JSON(http.StatusNotFound, dto.Error{Message: error.Error(), Cause: error.Cause(), StackTrace: error.StackTrace()})
	case *errapp.ConflictError:
		error := err.(*errapp.ConflictError)
		c.JSON(http.StatusConflict, dto.Error{Message: error.Error(), Cause: error.Cause(), StackTrace: error.StackTrace()})
	case *errapp.UnprocessableEntityError:
		error := err.(*errapp.UnprocessableEntityError)
		c.JSON(http.StatusUnprocessableEntity, dto.Error{Message: error.Error(), Cause: error.Cause(), StackTrace: error.StackTrace()})
	default:
		c.JSON(http.StatusInternalServerError, dto.Error{Message: err.(error).Error(), StackTrace: string(debug.Stack())})
	}
}
