package main

import (
	"fmt"
	"log"

	"github.com/evm-disassembler/internal/api"
	"github.com/evm-disassembler/internal/config"
	"github.com/evm-disassembler/internal/logger"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logg, err := logger.NewLogger(cfg.Logger.Level)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer logg.Sync()

	r := api.SetupRouter(logg)

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	logg.Info("Starting EVM Disassembler server",
		zap.String("address", addr),
		zap.String("log_level", cfg.Logger.Level),
	)

	if err := r.Run(addr); err != nil {
		logg.Fatal("Failed to run server", zap.Error(err))
	}
}
