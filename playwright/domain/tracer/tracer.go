//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock

package tracer

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type Tracer interface {
	Tracer() trace.Tracer
	Shutdown(context.Context) error
}
