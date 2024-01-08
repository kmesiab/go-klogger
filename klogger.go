package go_klogger

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

const DefaultLogLevel = logrus.DebugLevel

var (
	once          sync.Once
	globalLogger  *logrus.Logger
	defaultFields map[string]interface{}
)

type KLogger struct {
	Logger   *logrus.Logger
	Message  string                 `json:"message"`
	LogLevel logrus.Level           `json:"log_level"`
	Data     map[string]interface{} `json:"data"`
}

// InitializeGlobalLogger Optionally allows you to specify the logger level and formatter
// for the global logger.  If not called, the global logger will be initialized with default
// preferences.  If called multiple times, the global logger will not be overwritten.
//
// This function is thread safe.  It is safe to call this function from multiple goroutines.
func InitializeGlobalLogger(level logrus.Level, formatter logrus.Formatter) {

	once.Do(func() {
		globalLogger = logrus.New()
		globalLogger.SetLevel(level)
		globalLogger.SetFormatter(formatter)
	})
}

// SetLogger let's you replace the global logger.
func SetLogger(logger *logrus.Logger) {
	globalLogger = logger
}

// SetDefaultFields sets the default fields to be added to every log message.
// Typically, this will be properties like "app_name", "app_version", etc.
func SetDefaultFields(fields map[string]interface{}) {
	defaultFields = fields
}

// Logf creates a new logger with the given format and arguments.
func Logf(format string, a ...interface{}) *KLogger {

	// Set up a global logger with default preferences.  This is
	// only ever done once, whether here or by calling InitializeGlobalLogger
	// directly.
	InitializeGlobalLogger(DefaultLogLevel, &logrus.JSONFormatter{})

	// Pass back an instance of a KLogger with the global logger and default properties.
	k := &KLogger{
		Logger:  globalLogger,
		Message: fmt.Sprintf(format, a...),
		Data:    make(map[string]interface{}),
	}

	return k.AddData(defaultFields)
}

// SetLogLevel sets the log level of the logger.
func (l *KLogger) SetLogLevel(level logrus.Level) *KLogger {
	l.LogLevel = level
	l.Logger.SetLevel(level)

	return l
}

// Add adds a key-value pair to the logger's data.
func (l *KLogger) Add(key string, value interface{}) *KLogger {
	l.Data[key] = value

	return l
}

func (l *KLogger) Info() *KLogger {
	if l.LogLevel <= logrus.InfoLevel {
		l.Logger.WithFields(l.Data).Info(l.Message)
	}

	return l
}

func (l *KLogger) Warn() *KLogger {
	if l.LogLevel <= logrus.WarnLevel {
		l.Logger.WithFields(l.Data).Warn(l.Message)
	}

	return l
}

func (l *KLogger) Debug() *KLogger {
	if l.LogLevel <= logrus.DebugLevel {
		l.Logger.WithFields(l.Data).Debug(l.Message)
	}

	return l
}

func (l *KLogger) Error() *KLogger {
	if l.LogLevel <= logrus.ErrorLevel {
		l.Logger.WithFields(l.Data).Error(l.Message)
	}

	return l
}

func (l *KLogger) Fatal() *KLogger {
	if l.LogLevel <= logrus.FatalLevel {
		l.Logger.WithFields(l.Data).Fatal(l.Message)
	}

	return l
}

func (l *KLogger) Panic() *KLogger {
	if l.LogLevel <= logrus.PanicLevel {
		l.Logger.WithFields(l.Data).Panic(l.Message)
	}

	return l
}

func (l *KLogger) AddData(data map[string]interface{}) *KLogger {
	for k, v := range data {
		l.Data[k] = v
	}

	return l
}
