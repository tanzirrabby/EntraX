package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

func New(level string) zerolog.Logger {
	parsedLevel := zerolog.InfoLevel
	if lvl, err := zerolog.ParseLevel(strings.ToLower(level)); err == nil {
		parsedLevel = lvl
	}
	zerolog.SetGlobalLevel(parsedLevel)
	return zerolog.New(os.Stdout).With().Timestamp().Logger()
}
