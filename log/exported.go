package log

import (
	"fmt"
	stdLog "log"
	"os"
)

var defaultLogger FieldLogger

func init() {
	logger := stdLog.New(os.Stderr, "", stdLog.LstdFlags)
	stdLogger := NewStd(StdLevelDebug, logger)
	SetLogger(stdLogger)
}

func SetLogger(logger Interface) {
	defaultLogger = AdaptFieldLogger(logger)
}

func WithError(err error) *Entry {
	return defaultLogger.WithError(err)
}

func WithField(name string, value interface{}) *Entry {
	return defaultLogger.WithField(name, value)
}

func WithFields(fields map[string]interface{}) *Entry {
	return defaultLogger.WithFields(fields)
}

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...interface{}) {
	defaultLogger.Trace(args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// Panic logs a message at level Panic on the standard logger then panic with args.
func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

// Tracef logs a message at level Trace on the standard logger.
func Tracef(format string, args ...interface{}) {
	defaultLogger.Trace(fmt.Sprintf(format, args...))
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debug(fmt.Sprintf(format, args...))
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	defaultLogger.Info(fmt.Sprintf(format, args...))
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warn(fmt.Sprintf(format, args...))
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	defaultLogger.Error(fmt.Sprintf(format, args...))
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatal(fmt.Sprintf(format, args...))
}

// Panicf logs a message at level Panic on the standard logger then panic with args.
func Panicf(format string, args ...interface{}) {
	defaultLogger.Panic(fmt.Sprintf(format, args...))
}
