package log

import (
	"log/slog"
	"os"
	"todo/webservice/services/config"
)

type STDLogger struct {
	logger *slog.Logger
}

func NewSTDLogger(envConfig *config.ENVConfig) *STDLogger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     parseStringToLogLevel(envConfig.LogLevel),
	}))
	return &STDLogger{
		logger: logger,
	}
}

func (l *STDLogger) GetLogger() *slog.Logger {
	return l.logger
}

func (l *STDLogger) Info(message string) {
	l.logger.Info(message)
}

func (l *STDLogger) Err(e error) {
	l.logger.Error(e.Error())
}
