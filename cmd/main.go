package main

import (
	"Personal-Notes/internal/config"
	"Personal-Notes/internal/logging"
	"Personal-Notes/internal/repository/postgres"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

func main() {
	cfg := initConfig()
	logger := initLogger(cfg.AppEnv, cfg.LogFormat)

	db := initDBConnection(cfg, logger)
	defer func() {
		db.Close()
		logger.Info("db connection closed", logging.NewField("database", cfg.DBName))
	}()
}

func initConfig() *config.Config {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("fail[config]: ", err)
	}
	return cfg
}
func initLogger(appEnv string, logFormat string) *logging.ZapLogger {
	var cfg zap.Config

	if appEnv == "production" {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}

	if logFormat == "console" {
		cfg.Encoding = "console"
		cfg.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	} else {
		cfg.Encoding = "json"
		cfg.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	}

	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.MessageKey = "message"
	cfg.EncoderConfig.CallerKey = "caller"
	cfg.EncoderConfig.StacktraceKey = "stacktrace"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	cfg.InitialFields = map[string]interface{}{"service": "personal_notes_api"}

	logger, err := logging.NewZapLogger(cfg)
	if err != nil {
		log.Fatal("fail[logger]: ", err)
	}

	logger.Info("init[logger]: successfully initialized")
	return logger
}
func initDBConnection(cfg *config.Config, logger logging.Logger) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := postgres.NewPostgresDB(ctx, cfg)
	if err != nil {
		logger.Fatal("fail[db]: failed to initialize db connection", logging.NewField("error", err))
	}

	logger.Info("init[db]: successfully initialized db connection", logging.NewField("database", cfg.DBName))
	return db
}
