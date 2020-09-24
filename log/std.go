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
	// StdLevelPanic level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	StdLevelPanic StdLevel = iota
	// StdLevelFatal level. Logs and then calls `os.Exit(1)`.
	StdLevelFatal
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
	// StdLevelTrace level. Designates finer-grained informational events than the Debug.
	StdLevelTrace
)

func stdLevelText(level StdLevel) string {
	switch level {
	case StdLevelPanic:
		return "PANIC"
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
	case StdLevelTrace:
		return "TRACE"
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

func (s *Std) Panic(args ...interface{}) {
	s.print(StdLevelPanic, args...)
	panic(fmt.Sprintln(args...))
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

func (s *Std) Trace(args ...interface{}) {
	s.print(StdLevelTrace, args...)
}

func (s *Std) print(level StdLevel, args ...interface{}) {
	if level > s.level {
		return
	}
	text := stdLevelText(level) + " " + fmt.Sprintln(args...) // fmt.Sprintln makes sure spaces are always added between operands
	_ = s.logger.Output(3, text)
}
