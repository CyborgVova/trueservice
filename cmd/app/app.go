package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/cyborgvova/trueservice/internal/config"
)

const (
	envLocal   = "local"
	envDevelop = "development"
	envProduct = "production"
)

func main() {
	cfg := config.MustParse()

	logger := NewLogger(cfg.Env)

	logger.With(slog.String("env", cfg.Env))

	logger.Info("initializing server", slog.String("address", cfg.Host+":"+cfg.Port))
	logger.Debug("logger debug mode enabled")

	fmt.Println(cfg)
}

func NewLogger(env string) *slog.Logger {
	var logger *slog.Logger
	switch env {
	case envLocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDevelop:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProduct:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo.Level()}))
	}
	return logger
}
