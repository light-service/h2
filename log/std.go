package log

import (
	"fmt"
	stdLog "log"
	"os"
)

// StdLoggerLevel type
type StdLoggerLevel uint32

// These are the different logging levels.
const (
	// StdLoggerLevelFatal level. Logs and then calls `os.Exit(1)`.
	StdLoggerLevelFatal StdLoggerLevel = iota
	// StdLoggerLevelError level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	StdLoggerLevelError
	// StdLoggerLevelWarn level. Non-critical entries that deserve eyes.
	StdLoggerLevelWarn
	// StdLoggerLevelInfo level. General operational entries about what's going on inside the
	// application.
	StdLoggerLevelInfo
	// StdLoggerLevelDebug level. Usually only enabled when debugging. Very verbose logging.
	StdLoggerLevelDebug
)

func stdLevelText(level StdLoggerLevel) string {
	switch level {
	case StdLoggerLevelFatal:
		return "FATAL"
	case StdLoggerLevelError:
		return "ERROR"
	case StdLoggerLevelWarn:
		return "WARN"
	case StdLoggerLevelInfo:
		return "INFO"
	case StdLoggerLevelDebug:
		return "DEBUG"
	default:
		return "UNKNOWN"
	}
}

type StdLogger struct {
	level  StdLoggerLevel
	logger *stdLog.Logger
}

func AdaptStdLogger(level StdLoggerLevel, logger *stdLog.Logger) *StdLogger {
	return &StdLogger{
		level:  level,
		logger: logger,
	}
}

func (s *StdLogger) SetLevel(level StdLoggerLevel) {
	s.level = level
}

func (s *StdLogger) Fatal(args ...interface{}) {
	s.print(StdLoggerLevelFatal, args...)
	os.Exit(1)
}

func (s *StdLogger) Fatalf(format string, args ...interface{}) {
	s.printf(StdLoggerLevelFatal, format, args...)
	os.Exit(1)
}
func (s *StdLogger) Error(args ...interface{}) {
	s.print(StdLoggerLevelError, args...)
}

func (s *StdLogger) Errorf(format string, args ...interface{}) {
	s.printf(StdLoggerLevelError, format, args...)
}

func (s *StdLogger) Warn(args ...interface{}) {
	s.print(StdLoggerLevelWarn, args...)
}

func (s *StdLogger) Warnf(format string, args ...interface{}) {
	s.printf(StdLoggerLevelWarn, format, args...)
}

func (s *StdLogger) Info(args ...interface{}) {
	s.print(StdLoggerLevelInfo, args...)
}

func (s *StdLogger) Infof(format string, args ...interface{}) {
	s.printf(StdLoggerLevelInfo, format, args...)
}

func (s *StdLogger) Debug(args ...interface{}) {
	s.print(StdLoggerLevelDebug, args...)
}

func (s *StdLogger) Debugf(format string, args ...interface{}) {
	s.printf(StdLoggerLevelDebug, format, args...)
}

func (s *StdLogger) printf(level StdLoggerLevel, format string, args ...interface{}) {
	s.print(level, fmt.Sprintf(format, args))
}

func (s *StdLogger) print(level StdLoggerLevel, args ...interface{}) {
	if level > s.level {
		return
	}
	_ = s.logger.Output(3, stdLevelText(level)+" "+fmt.Sprint(args...))
}
