package api

import (
	"net/http"

	"github.com/devlongs/evm-disassembler/internal/disassembler"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DisassembleRequest struct {
	Bytecode string `json:"bytecode" binding:"required"`
}

func DisassembleHandler(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req DisassembleRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			logger.Error("Invalid request payload", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
			return
		}

		instructions, err := disassembler.DisassembleEVMBytecode(req.Bytecode)
		if err != nil {
			logger.Error("Failed to disassemble bytecode", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, instructions)
	}
}
