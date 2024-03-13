package goklogger

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// StringToLogrusLevel takes a string representation of a log level (e.g., "info", "debug") and returns
// the corresponding logrus.Level. It is designed to be flexible, accepting any case and trimming
// unnecessary whitespace from the input. This function supports all standard logrus levels:
// "trace", "debug", "info", "warn"/"warning", "error", "fatal", and "panic". If the input does not
// match any of these levels, it defaults to logrus.InfoLevel.
//
// Parameters:
//   - level: A string representing the desired log level. This parameter is case-insensitive and
//     spaces before and after the level name are ignored.
//
// Returns:
//   - A logrus.Level corresponding to the input string. If the input is unrecognized, the function
//     defaults to returning logrus.InfoLevel.
//
// Usage:
// This function is particularly useful for parsing log level configurations from environment
// variables or configuration files where the input may vary in case or include additional spaces.
// It ensures that the application can dynamically adjust logging verbosity based on runtime
// configurations.
//
// Example:
// logLevel := ParseLogLevel("  DeBug ")
// log.SetLevel(logLevel)
//
// Note:
// The function treats "warn" and "warning" as equivalent, mapping both to logrus.WarnLevel,
// aligning with logrus's own handling of these terms.
func StringToLogrusLevel(level string) logrus.Level {
	level = strings.ToLower(strings.TrimSpace(level))

	switch level {
	case "trace":
		return logrus.TraceLevel

	case "debug":
		return logrus.DebugLevel

	case "info":
		return logrus.InfoLevel

	case "warn", "warning": // logrus treats "warn" and "warning" as equivalent

		return logrus.WarnLevel
	case "error":

		return logrus.ErrorLevel
	case "fatal":

		return logrus.FatalLevel
	case "panic":

		return logrus.PanicLevel
	default:
		// It's a common practice to default to InfoLevel when an unknown level is provided
		return logrus.InfoLevel
	}
}

// LogLevel represents logging levels by their severity.
type LogLevel int

const (
	TraceLevel LogLevel = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)

// StringToLogLevel interprets a string as a logging level and converts it to the corresponding
// LogLevel constant defined in this package. The function is designed to be case-insensitive
// and ignores any leading or trailing whitespace in the input string. Recognized levels include
// "trace", "debug", "info", "warn"/"warning", "error", "fatal", and "panic". If the input string
// does not match any of these recognized levels, the function defaults to returning InfoLevel.
//
// This flexibility allows for dynamic adjustment of logging verbosity based on runtime
// configuration, facilitating easy integration with various configuration sources such as
// environment variables or configuration files.
//
// Parameters:
// - level: A string representing the logging level to parse.
//
// Returns:
//   - The LogLevel constant corresponding to the provided string. Defaults to InfoLevel if the
//     input string is unrecognized.
//
// Example usage:
// logLevel := logger.ParseLogLevel("error")
// logger.SetLevel(logLevel)
//
// The function ensures compatibility with various input formats by treating "warn" and "warning"
// as equivalent, aligning with common logging conventions.
func StringToLogLevel(level string) LogLevel {
	level = strings.ToLower(strings.TrimSpace(level))

	switch level {
	case "trace":

		return TraceLevel
	case "debug":

		return DebugLevel
	case "info":

		return InfoLevel
	case "warn", "warning": // Support both "warn" and "warning" for compatibility

		return WarnLevel
	case "error":

		return ErrorLevel
	case "fatal":

		return FatalLevel
	case "panic":

		return PanicLevel
	default:

		return InfoLevel // Default to InfoLevel for unrecognized strings
	}
}
