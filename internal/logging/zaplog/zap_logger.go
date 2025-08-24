package zaplog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"Personal-Notes/internal/logging"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(cfg zap.Config) (*ZapLogger, error) {
	zapLogger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return &ZapLogger{logger: zapLogger}, nil
}

func (l *ZapLogger) Info(msg string, fields ...logging.Field) {
	zapFields := toZapFields(fields)
	l.logger.Info(msg, zapFields...)
}

func (l *ZapLogger) Debug(msg string, fields ...logging.Field) {
	zapFields := toZapFields(fields)
	l.logger.Debug(msg, zapFields...)
}

func (l *ZapLogger) Warn(msg string, fields ...logging.Field) {
	zapFields := toZapFields(fields)
	l.logger.Warn(msg, zapFields...)
}

func (l *ZapLogger) Error(msg string, fields ...logging.Field) {
	zapFields := toZapFields(fields)
	l.logger.Error(msg, zapFields...)
}

func (l *ZapLogger) Fatal(msg string, fields ...logging.Field) {
	zapFields := toZapFields(fields)
	l.logger.Fatal(msg, zapFields...)
}

func toZapFields(fields []logging.Field) []zapcore.Field {
	zapFields := make([]zapcore.Field, 0, len(fields))
	for _, f := range fields {
		zapFields = append(zapFields, zap.Any(f.Key, f.Value))
	}
	return zapFields
}
