package middleware

import (
	httperr "ortisan-broker/go-commons/infrastructure/http/error"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httperr.HandleError(c, err.(error))
			}
		}()
		c.Next()
	}
}
