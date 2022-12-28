package middleware

import (
	"net/http"

	"github.com/Pranc1ngPegasus/golang-lab/playwright/domain/logger"
)

func Recover(
	logger logger.Logger,
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			defer func() {
				if p := recover(); p != nil {
					err, ok := p.(error)
					if ok {
						logger.Error(ctx, "panic occurred", logger.Field("err", err))
					}

					w.WriteHeader(http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
