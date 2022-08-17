package handler

import (
	"net/http"
)

func NewHandler(
	logging func(http.Handler) http.Handler,
	healthcheck Healthcheck,
) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/healthcheck", logging(healthcheck))

	return mux
}
