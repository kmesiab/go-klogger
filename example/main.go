package main

import (
	"github.com/sirupsen/logrus"

	log "github.com/kmesiab/go-klogger"
)

func main() {

	// Set up the default fields to be added to every log message
	defaultFields := map[string]interface{}{
		"foo": "bar",
		"baz": "qux",
	}

	log.SetDefaultFields(defaultFields)

	// Set up the global logger with default preferences
	log.InitializeGlobalLogger(logrus.WarnLevel, &logrus.TextFormatter{
		DisableTimestamp: true,
	})

	// Log a message using the `Logf()` function
	log.Logf("Hello %s", "World").Warn()

	// Add extra fields to the log message
	log.Logf("Hello %s", "World").
		Add("foo", "bar").
		Info()

	// Replace the global logger
	log.SetLogger(logrus.New())
	log.Logf("Hello from a new logger!").Info()

}
