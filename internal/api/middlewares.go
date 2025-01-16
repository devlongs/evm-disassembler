package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ExampleMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// do something before request
		logger.Debug("Before handling request")

		c.Next()

		// do something after request
		logger.Debug("After handling request")
	}
}
