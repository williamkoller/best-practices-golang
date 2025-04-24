package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strings"
)

func NewLogger() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	levelStr := strings.ToLower(os.Getenv("LOG_LEVEL"))
	level, err := zerolog.ParseLevel(levelStr)
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)

	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}

	isDev := os.Getenv("APP_ENV") == "local" || os.Getenv("APP_ENV") == "development"

	var baseLogger zerolog.Logger
	if isDev {
		baseLogger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().
			Timestamp().
			Str("app", os.Getenv("APP_NAME")).
			Str("env", os.Getenv("APP_ENV")).
			Caller().
			Logger()
	} else {
		baseLogger = zerolog.New(os.Stderr).With().
			Timestamp().
			Str("app", os.Getenv("APP_NAME")).
			Str("env", os.Getenv("APP_ENV")).
			Caller().
			Logger()
	}

	log.Logger = baseLogger
	return baseLogger
}
