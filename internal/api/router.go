package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupRouter(logger *zap.Logger) *gin.Engine {
	r := gin.Default()

	// r.Use(gin.Recovery())

	api := r.Group("/api")
	{
		api.POST("/disassemble", DisassembleHandler(logger))
	}

	return r
}
