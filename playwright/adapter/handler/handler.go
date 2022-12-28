package handler

import (
	"net/http"

	"github.com/Pranc1ngPegasus/golang-lab/playwright/adapter/handler/middleware"
	"github.com/Pranc1ngPegasus/golang-lab/playwright/domain/logger"
	"github.com/Pranc1ngPegasus/golang-lab/playwright/domain/usecase"
	"github.com/google/wire"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

var _ http.Handler = (*Handler)(nil)

var NewHandlerSet = wire.NewSet(
	wire.Bind(new(http.Handler), new(*Handler)),
	NewHandler,
)

type Handler struct {
	logger logger.Logger
	router http.Handler

	playwrightUsecase usecase.Playwright
}

func NewHandler(
	logger logger.Logger,
	playwrightUsecase usecase.Playwright,
) *Handler {
	mux := http.NewServeMux()

	router := otelhttp.NewHandler(
		middleware.Chain(mux,
			middleware.Recover(logger),
			middleware.Logger(logger),
		),
		"playwright",
		otelhttp.WithSpanNameFormatter(
			func(operation string, req *http.Request) string {
				return req.URL.Path
			},
		),
	)

	h := &Handler{
		logger:            logger,
		router:            router,
		playwrightUsecase: playwrightUsecase,
	}

	mux.HandleFunc("/healthcheck", h.healthcheck)
	mux.HandleFunc("/scan", h.scan)

	return h
}

func (h *Handler) healthcheck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	_, err := w.Write([]byte("ok"))
	if err != nil {
		h.logger.Error(ctx, "failed to write healthcheck response", h.logger.Field("err", err))
	}
}

func (h *Handler) scan(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	_, err := h.playwrightUsecase.Do(ctx, usecase.PlaywrightInput{
		URL: "https://www.google.com/",
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
