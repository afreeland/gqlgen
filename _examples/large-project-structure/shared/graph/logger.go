package graph

import "io"

// Logger is a generic interface for logging with various levels and namespaces.
type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
}

// LoggerConfig holds configuration for creating loggers.
type LoggerConfig struct {
	Output io.Writer
}
