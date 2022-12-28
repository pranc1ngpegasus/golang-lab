package server

import (
	"context"
	"net/http"
	"time"

	"github.com/Pranc1ngPegasus/golang-lab/playwright/domain/configuration"
	"github.com/Pranc1ngPegasus/golang-lab/playwright/domain/logger"
)

func NewServer(
	ctx context.Context,
	logger logger.Logger,
	cfg configuration.Configuration,
	handler http.Handler,
) *http.Server {
	config := cfg.Server()

	logger.Info(ctx, "listen on", logger.Field("port", config.Port))

	return &http.Server{
		Addr:              ":" + config.Port,
		Handler:           handler,
		ReadHeaderTimeout: 30 * time.Second,
	}
}
