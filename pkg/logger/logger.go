package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	EnvVarHumanFriendlyLogs = "HUMAN_FRIENDLY_LOGS"
)

// AllowedLevels is the list of allowed log levels
var AllowedLevels = []string{"debug", "info", "warn", "error", "fatal", "panic", "disabled", "trace"}

// CallerHook implements zerolog.Hook interface.
type CallerHook struct{}

// Run adds additional context
func (h CallerHook) Run(e *zerolog.Event, _ zerolog.Level, _ string) {
	if _, file, line, ok := runtime.Caller(3); ok {
		e.Str("file", fmt.Sprintf("%s:%d", path.Base(file), line))
	}
}

func newLogger(component string) zerolog.Logger {
	return log.With().Str("component", component).Logger().Hook(CallerHook{})
}

// New creates a new zerolog.Logger
func New(component string) zerolog.Logger {
	if os.Getenv(EnvVarHumanFriendlyLogs) == "true" {
		return NewPretty(component)
	}
	return newLogger(component)
}

// NewPretty creates a new zerolog.Logger, which emits human-readable log messages
func NewPretty(component string) zerolog.Logger {
	return newLogger(component).Output(zerolog.ConsoleWriter{Out: os.Stdout})
}

// SetLogLevel sets the global logger level
func SetLogLevel(verbosity string) error {
	switch strings.ToLower(verbosity) {
	// DebugLevel defines debug log level.
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)

	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)

	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)

	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)

	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)

	case "disabled":
		zerolog.SetGlobalLevel(zerolog.Disabled)

	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)

	default:
		return fmt.Errorf("invalid log level '%s' specified; please specify one of %v", verbosity, AllowedLevels)
	}
	return nil
}
