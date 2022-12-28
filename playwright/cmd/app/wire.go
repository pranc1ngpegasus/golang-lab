//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"net/http"

	"github.com/Pranc1ngPegasus/golang-lab/playwright/adapter/handler"
	"github.com/Pranc1ngPegasus/golang-lab/playwright/adapter/server"
	domainlogger "github.com/Pranc1ngPegasus/golang-lab/playwright/domain/logger"
	domaintracer "github.com/Pranc1ngPegasus/golang-lab/playwright/domain/tracer"
	"github.com/Pranc1ngPegasus/golang-lab/playwright/infra/configuration"
	"github.com/Pranc1ngPegasus/golang-lab/playwright/infra/logger"
	"github.com/Pranc1ngPegasus/golang-lab/playwright/infra/tracer"
	"github.com/google/wire"
)

type app struct {
	ctx    context.Context
	logger domainlogger.Logger
	tracer domaintracer.Tracer
	server *http.Server
}

func initialize() (*app, error) {
	wire.Build(
		context.Background,
		configuration.NewConfigurationSet,
		logger.NewLoggerSet,
		tracer.NewTracerSet,
		handler.NewHandlerSet,
		server.NewServer,
		wire.Struct(new(app), "*"),
	)

	return nil, nil
}
