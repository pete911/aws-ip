package internal

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

var logLevels = map[string]slog.Level{"debug": slog.LevelDebug, "info": slog.LevelInfo, "warn": slog.LevelWarn, "error": slog.LevelError}

func NewLogger(logLevel string) (*slog.Logger, error) {
	if level, ok := logLevels[strings.ToLower(logLevel)]; ok {
		opts := &slog.HandlerOptions{Level: level}
		return slog.New(slog.NewTextHandler(os.Stderr, opts)), nil
	}
	return nil, fmt.Errorf("invalid log level %s", logLevel)
}

func NewNoopLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}
