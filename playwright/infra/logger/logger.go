package logger

import (
	"context"
	"fmt"

	"github.com/Pranc1ngPegasus/golang-lab/playwright/domain/configuration"
	domain "github.com/Pranc1ngPegasus/golang-lab/playwright/domain/logger"
	"github.com/google/wire"
	"github.com/samber/lo"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ domain.Logger = (*Logger)(nil)

var NewLoggerSet = wire.NewSet(
	wire.Bind(new(domain.Logger), new(*Logger)),
	NewLogger,
)

type Logger struct {
	config *configuration.Common
	logger *otelzap.Logger
}

func NewLogger(
	cfg configuration.Configuration,
) (*Logger, error) {
	config := cfg.Common()

	zapconfig := zap.NewProductionConfig()
	zapconfig.EncoderConfig = encoderConfig()

	if config.Debug {
		zapconfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	zaplogger, err := zapconfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	logger := otelzap.New(zaplogger)

	return &Logger{
		config: config,
		logger: logger,
	}, nil
}

func encoderConfig() zapcore.EncoderConfig {
	cfg := zap.NewProductionEncoderConfig()
	cfg.LevelKey = "severity"
	cfg.EncodeLevel = EncodeLevel
	cfg.TimeKey = "time"
	cfg.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	return cfg
}

// Cloud Loggingのseverityと合わせる.
var logLevelSeverity = map[zapcore.Level]string{
	zapcore.DebugLevel:  "DEBUG",
	zapcore.InfoLevel:   "INFO",
	zapcore.WarnLevel:   "WARNING",
	zapcore.ErrorLevel:  "ERROR",
	zapcore.DPanicLevel: "CRITICAL",
	zapcore.PanicLevel:  "ALERT",
	zapcore.FatalLevel:  "EMERGENCY",
}

func EncodeLevel(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(logLevelSeverity[l])
}

func (l *Logger) Field(key string, message interface{}) domain.Field {
	return domain.Field{
		Key:       key,
		Interface: message,
	}
}

func (l *Logger) field(field domain.Field) zap.Field {
	switch i := field.Interface.(type) {
	case error:
		return zap.Error(i)
	case string:
		return zap.String(field.Key, i)
	case int:
		return zap.Int(field.Key, i)
	case bool:
		return zap.Bool(field.Key, i)
	default:
		return zap.Any(field.Key, i)
	}
}

func (l *Logger) traceAttributes(ctx context.Context) []domain.Field {
	fields := []domain.Field{}

	span := trace.SpanContextFromContext(ctx)
	if span.TraceID().IsValid() {
		fields = append(
			fields,
			l.Field("logging.googleapis.com/trace", fmt.Sprintf("projects/%s/traces/%s", l.config.GCPProjectID, span.TraceID())),
			l.Field("logging.googleapis.com/spanId", span.SpanID().String()),
			l.Field("logging.googleapis.com/trace_sampled", span.IsSampled()),
		)
	}

	return fields
}

func (l *Logger) Debug(ctx context.Context, message string, fields ...domain.Field) {
	fields = append(fields, l.traceAttributes(ctx)...)
	zapfields := lo.Map(fields, func(field domain.Field, _ int) zap.Field {
		return l.field(field)
	})

	l.logger.Ctx(ctx).Debug(message, zapfields...)
}

func (l *Logger) Info(ctx context.Context, message string, fields ...domain.Field) {
	fields = append(fields, l.traceAttributes(ctx)...)
	zapfields := lo.Map(fields, func(field domain.Field, _ int) zap.Field {
		return l.field(field)
	})

	l.logger.Ctx(ctx).Info(message, zapfields...)
}

func (l *Logger) Error(ctx context.Context, message string, fields ...domain.Field) {
	fields = append(fields, l.traceAttributes(ctx)...)
	zapfields := lo.Map(fields, func(field domain.Field, _ int) zap.Field {
		return l.field(field)
	})

	l.logger.Ctx(ctx).Error(message, zapfields...)
}
