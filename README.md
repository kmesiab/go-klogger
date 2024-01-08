
# go_klogger - A Go Logging Package

![Golang](https://img.shields.io/badge/Go-00add8.svg?labelColor=171e21&style=for-the-badge&logo=go)

![Build](https://github.com/kmesiab/go-klogger/actions/workflows/go.yml/badge.svg)

`go_klogger` is a flexible logging package for Go, built on top of the popular `logrus` package. 
It provides easy-to-use functions for logging at different levels with support for customizable 
fields and global configuration.

## Installation

To use `go_klogger`, first install it using `go get`:

```bash
go get github.com/kmesiab/go-klogger
```

## Usage

Here's a quick guide to get you started:

### Setting Up Default Fields

You can set up default fields which will be added to every log message. Typically,
these are properties like "app_name", "app_version", etc.

```go
log.SetDefaultFields(map[string]interface{}{
"app_name":    "MyApplication",
"app_version": "1.0.0",
})
```

### Initializing the Global Logger

You can initialize the global logger with your preferred log level and formatter.
This step is optional; if not called, the global logger will be initialized with
default preferences.

```go
log.InitializeGlobalLogger(logrus.InfoLevel, &logrus.JSONFormatter{})
```

### Logging Messages

Use `Logf` to create a new logger instance with a formatted message, and
then chain it with one of the logging methods
(`Info`, `Warn`, `Debug`, `Error`, `Fatal`, `Panic`) based on the required 
severity.

```go
log.Logf("This is an %s message", "info").Info()
log.Logf("This is a warning!").Warn()
```

### Adding Additional Data

You can add more key-value pairs to your log data easily.

```go
log.Logf("User login attempt").Add("username", "johndoe").Info()
```

### Replacing the Global Logger

If needed, you can replace the global logger instance with a new one.

```go
newLogger := logrus.New()
newLogger.SetFormatter(&logrus.TextFormatter{})
log.SetLogger(newLogger)
log.Logf("Logging with a new logger instance").Info()
```
## Example Usage

Here is an example of how to use `go_klogger` in your application:

```go
package main

import (
    "github.com/sirupsen/logrus"
    log "github.com/kmesiab/go-klogger"
)

func main() {
	
    // Set up default fields
    log.SetDefaultFields(map[string]interface{}{
        "foo": "bar",
        "baz": "qux",
    })

    // Initialize the global logger
    log.InitializeGlobalLogger(logrus.WarnLevel, &logrus.TextFormatter{
        DisableTimestamp: true,
    })

    // Log a message
    log.Logf("Hello %s", "World").Warn()

    // Add extra fields and log at a different level
    log.Logf("Hello %s", "World").
        Add("additional", "info").
        Info()

    // Replace the global logger and log a message
    log.SetLogger(logrus.New())
    log.Logf("Hello from a new logger!").Info()
}
```

This example demonstrates setting default fields, initializing the global logger
with specific preferences, logging messages at different levels, adding additional
data to logs, and replacing the global logger.

---
## Contributing

We welcome contributions! If you'd like to contribute to `go_klogger`, please follow
these steps:

1. Fork the repository on GitHub.
2. Clone your forked repository to your local machine.
3. Create a new branch for your feature or bug fix.
4. Make your changes and commit them to your branch.
5. Push your changes to your GitHub fork.
6. Open a Pull Request against the `main` branch of `go_klogger` repository.

For more details, check out GitHub's guide on [forking](https://docs.github.com/en/get-started/quickstart/fork-a-repo)
and [creating a pull request](https://docs.github.com/en/get-started/quickstart/contributing-to-projects).
