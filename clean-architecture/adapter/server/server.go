package server

import (
	"net/http"
	"time"

	"github.com/Pranc1ngPegasus/golang-lab/clean-architecture/adapter/configuration"
)

func NewServer(
	config configuration.Server,
	handler http.Handler,
) *http.Server {
	return &http.Server{
		Addr:              ":" + config.Port,
		Handler:           handler,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}
}
