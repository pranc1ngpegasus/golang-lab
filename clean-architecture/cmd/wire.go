//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/Pranc1ngPegasus/golang-lab/clean-architecture/adapter/configuration"
	"github.com/Pranc1ngPegasus/golang-lab/clean-architecture/adapter/handler"
	"github.com/Pranc1ngPegasus/golang-lab/clean-architecture/adapter/handler/middleware"
	"github.com/Pranc1ngPegasus/golang-lab/clean-architecture/adapter/logger"
	"github.com/Pranc1ngPegasus/golang-lab/clean-architecture/adapter/server"
	domainlogger "github.com/Pranc1ngPegasus/golang-lab/clean-architecture/domain/logger"
	"github.com/google/wire"
)

type app struct {
	logger domainlogger.Logger
	server *http.Server
}

func initialize() (*app, error) {
	wire.Build(
		logger.NewLoggerSet,

		configuration.Get,
		wire.FieldsOf(new(configuration.Config), "Server"),

		middleware.NewLogging,

		handler.NewHealthcheck,
		handler.NewHandler,

		server.NewServer,

		wire.Struct(new(app), "*"),
	)

	return nil, nil
}
