package log

import (
	"log/slog"
)

func parseStringToLogLevel(s string) slog.Level {
	level := slog.LevelError
	level.UnmarshalText([]byte(s))
	return level
}
