//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock

package logger

import (
	"context"
)

type Logger interface {
	Field(key string, message interface{}) Field
	Debug(ctx context.Context, key string, fields ...Field)
	Info(ctx context.Context, key string, fields ...Field)
	Error(ctx context.Context, key string, fields ...Field)
}

type (
	Field struct {
		Key       string
		Interface interface{}
	}
)
