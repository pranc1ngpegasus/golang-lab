package logger

import (
	"fmt"

	domain "github.com/Pranc1ngPegasus/golang-lab/clean-architecture/domain/logger"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var _ domain.Logger = (*Logger)(nil)

var NewLoggerSet = wire.NewSet(
	wire.Bind(new(domain.Logger), new(*Logger)),
	NewLogger,
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger() (*Logger, error) {
	log, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	return &Logger{
		logger: log,
	}, nil
}

func (l *Logger) Info(message string, args map[string]interface{}) {
	fields := []zap.Field{}
	for k, v := range args {
		fields = append(fields, zap.Any(k, v))
	}

	l.logger.Info(message, fields...)
}

func (l *Logger) Error(message string, err error) {
	l.logger.Error(message, zap.Error(err))
}
