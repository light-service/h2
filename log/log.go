package log

import (
	"fmt"
	"strings"
)

type Interface interface {
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
	Trace(args ...interface{})
}

type Logger interface {
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
}

type logger struct {
	Interface
}

func AdaptLogger(l Interface) Logger {
	if _, ok := l.(*Entry); ok {
		panic("logger must not be *Entry")
	}
	if f, ok := l.(Logger); ok {
		return f
	}

	return &logger{l}
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.Panic(fmt.Sprintf(format, args...))
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.Fatal(fmt.Sprintf(format, args...))
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.Warn(fmt.Sprintf(format, args...))
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.Debug(fmt.Sprintf(format, args...))
}

func (l *logger) Tracef(format string, args ...interface{}) {
	l.Trace(fmt.Sprintf(format, args...))
}

type Fields map[string]interface{}

func (f Fields) Append(another Fields) {
	if another == nil {
		return
	}
	for k, v := range another {
		f[k] = v
	}
}

func (f Fields) String() string {
	pairs := make([]string, 0, len(f))
	for k, v := range f {
		pair := fmt.Sprintf("%s=%v", k, v)
		pairs = append(pairs, pair)
	}

	return "{ " + strings.Join(pairs, ", ") + " }"
}

func SplitArgs(mixedArgs []interface{}) (fields Fields, args []interface{}) {
	fields = make(Fields)
	for _, m := range mixedArgs {
		switch m.(type) {
		case Fields:
			fields.Append(m.(Fields))
		default:
			args = append(args, m)
		}
	}
	return
}

type Entry struct {
	logger Logger
	fields Fields
}

func newEntry(logger Logger, fields Fields) *Entry {
	return &Entry{logger, fields}
}

func (e *Entry) Panic(args ...interface{}) {
	e.logger.Panic(e.args(args)...)
}

func (e *Entry) Panicf(format string, args ...interface{}) {
	e.logger.Panic(e.argsf(format, args...)...)
}

func (e *Entry) Fatal(args ...interface{}) {
	e.logger.Fatal(e.args(args)...)
}

func (e *Entry) Fatalf(format string, args ...interface{}) {
	e.logger.Fatal(e.argsf(format, args...)...)
}

func (e *Entry) Error(args ...interface{}) {
	e.logger.Error(e.args(args)...)
}

func (e *Entry) Errorf(format string, args ...interface{}) {
	e.logger.Error(e.argsf(format, args...)...)
}

func (e *Entry) Warn(args ...interface{}) {
	e.logger.Warn(e.args(args)...)
}

func (e *Entry) Warnf(format string, args ...interface{}) {
	e.logger.Warn(e.argsf(format, args...)...)
}

func (e *Entry) Info(args ...interface{}) {
	e.logger.Info(e.args(args)...)
}

func (e *Entry) Infof(format string, args ...interface{}) {
	e.logger.Info(e.argsf(format, args...)...)
}

func (e *Entry) Debug(args ...interface{}) {
	e.logger.Debug(e.args(args)...)
}

func (e *Entry) Debugf(format string, args ...interface{}) {
	e.logger.Debug(e.argsf(format, args...)...)
}

func (e *Entry) Trace(args ...interface{}) {
	e.logger.Trace(e.args(args)...)
}

func (e *Entry) Tracef(format string, args ...interface{}) {
	e.logger.Trace(e.argsf(format, args...)...)
}

func (e *Entry) WithError(err error) *Entry {
	e.fields[errorFieldName] = err
	return e
}

func (e *Entry) WithField(name string, value interface{}) *Entry {
	e.fields[name] = value
	return e
}

func (e *Entry) WithFields(fields map[string]interface{}) *Entry {
	e.fields.Append(fields)
	return e
}

func (e *Entry) args(args ...interface{}) []interface{} {
	return append(args, e.fields)
}

func (e *Entry) argsf(format string, args ...interface{}) []interface{} {
	arg0 := fmt.Sprintf(format, args...)
	return []interface{}{arg0, e.fields}
}

const errorFieldName = "error"

type FieldLogger interface {
	Logger
	WithError(err error) *Entry
	WithField(name string, value interface{}) *Entry
	WithFields(fields map[string]interface{}) *Entry
}

type fieldLogger struct {
	Logger
}

func AdaptFieldLogger(l Interface) FieldLogger {
	if _, ok := l.(*Entry); ok {
		panic("logger must not be *Entry")
	}

	if f, ok := l.(FieldLogger); ok {
		return f
	} else if f, ok := l.(Logger); ok {
		return &fieldLogger{f}
	} else {
		return &fieldLogger{AdaptLogger(l)}
	}
}

func (f *fieldLogger) WithError(err error) *Entry {
	return newEntry(f.Logger, Fields{errorFieldName: err})
}

func (f *fieldLogger) WithField(name string, value interface{}) *Entry {
	return newEntry(f.Logger, Fields{name: value})
}

func (f *fieldLogger) WithFields(fields map[string]interface{}) *Entry {
	return newEntry(f.Logger, fields)
}
