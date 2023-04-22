package middleware

import (
	httpErr "ortisan-broker/go-commons/infrastructure/http/error"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpErr.HandleError(c, err.(error))
			}
		}()
		c.Next()
	}
}
