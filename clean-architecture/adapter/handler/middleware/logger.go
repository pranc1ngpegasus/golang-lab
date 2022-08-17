package middleware

import (
	"net/http"
	"time"

	domainlogger "github.com/Pranc1ngPegasus/golang-lab/clean-architecture/domain/logger"
)

func NewLogging(logger domainlogger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t1 := time.Now()

			defer func() {
				logger.Info("Served",
					map[string]interface{}{
						"proto":    r.Proto,
						"method":   r.Method,
						"path":     r.URL.Path,
						"duration": time.Since(t1).String(),
					},
				)
			}()

			next.ServeHTTP(w, r)
		})
	}
}
