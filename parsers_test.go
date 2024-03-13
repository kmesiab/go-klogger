package goklogger

import (
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestStringToLogrusLevel(t *testing.T) {
	tests := []struct {
		name          string
		level         string
		expectedLevel logrus.Level
	}{
		{name: "trace level", level: "trace", expectedLevel: logrus.TraceLevel},
		{name: "debug level", level: "debug", expectedLevel: logrus.DebugLevel},
		{name: "info level", level: "info", expectedLevel: logrus.InfoLevel},
		{name: "warn level", level: "warn", expectedLevel: logrus.WarnLevel},
		{name: "warning level", level: "warning", expectedLevel: logrus.WarnLevel}, // Test equivalent to "warn"
		{name: "error level", level: "error", expectedLevel: logrus.ErrorLevel},
		{name: "fatal level", level: "fatal", expectedLevel: logrus.FatalLevel},
		{name: "panic level", level: "panic", expectedLevel: logrus.PanicLevel},
		{name: "default to info level", level: "unknown", expectedLevel: logrus.InfoLevel},
		{name: "case insensitivity", level: "DeBuG", expectedLevel: logrus.DebugLevel},
		{name: "leading/trailing spaces", level: "  info  ", expectedLevel: logrus.InfoLevel},

		// Negative tests
		{name: "numeric level", level: "123", expectedLevel: logrus.InfoLevel},
		{name: "special characters", level: "@#&*", expectedLevel: logrus.InfoLevel},
		{name: "empty string", level: "", expectedLevel: logrus.InfoLevel},
		{name: "whitespace only", level: "   ", expectedLevel: logrus.InfoLevel},
		{name: "mixed characters", level: "info123", expectedLevel: logrus.InfoLevel},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualLevel := StringToLogrusLevel(tt.level)
			if actualLevel != tt.expectedLevel {
				t.Errorf("ParseLogLevel(%s) = %v, want %v", strings.TrimSpace(tt.level), actualLevel, tt.expectedLevel)
			}
		})
	}
}

func TestStringToLogLevel(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedLevel LogLevel
	}{
		{"Trace level", "trace", TraceLevel},
		{"Debug level", "debug", DebugLevel},
		{"Info level", "info", InfoLevel},
		{"Warn level", "warn", WarnLevel},
		{"Warning level", "warning", WarnLevel}, // Test equivalent to "warn"
		{"Error level", "error", ErrorLevel},
		{"Fatal level", "fatal", FatalLevel},
		{"Panic level", "panic", PanicLevel},
		{"Default to Info level", "unknown", InfoLevel},
		{"Case insensitivity", "DeBuG", DebugLevel},
		{"Leading/trailing spaces", "  info  ", InfoLevel},
		// Negative tests
		{"Numeric level", "123", InfoLevel},
		{"Special characters", "@#&*", InfoLevel},
		{"Empty string", "", InfoLevel},
		{"Whitespace only", "   ", InfoLevel},
		{"Mixed characters", "info123", InfoLevel},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualLevel := StringToLogLevel(tt.input)
			if actualLevel != tt.expectedLevel {
				t.Errorf("ParseLogLevel(%q) = %v, want %v", strings.TrimSpace(tt.input), actualLevel, tt.expectedLevel)
			}
		})
	}
}
