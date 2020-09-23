package log

import (
	"fmt"
	stdLog "log"
	"os"
)

// StdLevel type
type StdLevel uint32

// These are the different logging levels.
const (
	// StdLevelFatal level. Logs and then calls `os.Exit(1)`.
	StdLevelFatal StdLevel = iota
	// StdLevelError level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	StdLevelError
	// StdLevelWarn level. Non-critical entries that deserve eyes.
	StdLevelWarn
	// StdLevelInfo level. General operational entries about what's going on inside the
	// application.
	StdLevelInfo
	// StdLevelDebug level. Usually only enabled when debugging. Very verbose logging.
	StdLevelDebug
)

func stdLevelText(level StdLevel) string {
	switch level {
	case StdLevelFatal:
		return "FATAL"
	case StdLevelError:
		return "ERROR"
	case StdLevelWarn:
		return "WARN"
	case StdLevelInfo:
		return "INFO"
	case StdLevelDebug:
		return "DEBUG"
	default:
		return "UNKNOWN"
	}
}

type Std struct {
	level  StdLevel
	logger *stdLog.Logger
}

func NewStd(level StdLevel, logger *stdLog.Logger) *Std {
	return &Std{
		level:  level,
		logger: logger,
	}
}

func (s *Std) SetLevel(level StdLevel) {
	s.level = level
}

func (s *Std) Fatal(args ...interface{}) {
	s.print(StdLevelFatal, args...)
	os.Exit(1)
}

func (s *Std) Error(args ...interface{}) {
	s.print(StdLevelError, args...)
}

func (s *Std) Warn(args ...interface{}) {
	s.print(StdLevelWarn, args...)
}

func (s *Std) Info(args ...interface{}) {
	s.print(StdLevelInfo, args...)
}

func (s *Std) Debug(args ...interface{}) {
	s.print(StdLevelDebug, args...)
}

func (s *Std) print(level StdLevel, args ...interface{}) {
	if level > s.level {
		return
	}
	_ = s.logger.Output(3, stdLevelText(level)+" "+fmt.Sprintln(args...))
}
