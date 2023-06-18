package error

import (
	"net/http"
	"ortisan-broker/go-commons/adapter/dto"
	errapp "ortisan-broker/go-commons/error"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, errParam error) {
	switch errParam.(type) {
	case *errapp.BadArgumentError:
		err := errParam.(*errapp.BadArgumentError)
		c.JSON(http.StatusBadRequest, dto.Error{Message: err.Error(), Cause: err.Cause(), StackTrace: err.StackTrace()})
	case *errapp.AuthError:
		err := errParam.(*errapp.AuthError)
		c.JSON(http.StatusUnauthorized, dto.Error{Message: err.Error(), Cause: err.Cause(), StackTrace: err.StackTrace()})
	case *errapp.NotFoundError:
		err := errParam.(*errapp.NotFoundError)
		c.JSON(http.StatusNotFound, dto.Error{Message: err.Error(), Cause: err.Cause(), StackTrace: err.StackTrace()})
	case *errapp.ConflictError:
		err := errParam.(*errapp.ConflictError)
		c.JSON(http.StatusConflict, dto.Error{Message: err.Error(), Cause: err.Cause(), StackTrace: err.StackTrace()})
	case *errapp.UnprocessableEntityError:
		err := errParam.(*errapp.UnprocessableEntityError)
		c.JSON(http.StatusUnprocessableEntity, dto.Error{Message: err.Error(), Cause: err.Cause(), StackTrace: err.StackTrace()})
	default:
		c.JSON(http.StatusInternalServerError, dto.Error{Message: errParam.(error).Error(), StackTrace: string(debug.Stack())})
	}
}
