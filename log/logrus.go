package log

import "github.com/sirupsen/logrus"

type Logrus struct {
	logger *logrus.Logger
}

func NewLogrus(l *logrus.Logger) *Logrus {
	return &Logrus{
		logger: l,
	}
}

func (l *Logrus) Panic(mixedArgs ...interface{}) {
	fields, args := l.fields(mixedArgs)
	l.logger.WithFields(fields).Panic(args...)
}

func (l *Logrus) Fatal(mixedArgs ...interface{}) {
	fields, args := l.fields(mixedArgs)
	l.logger.WithFields(fields).Fatal(args...)
}

func (l *Logrus) Error(mixedArgs ...interface{}) {
	fields, args := l.fields(mixedArgs)
	l.logger.WithFields(fields).Error(args...)
}

func (l *Logrus) Warn(mixedArgs ...interface{}) {
	fields, args := l.fields(mixedArgs)
	l.logger.WithFields(fields).Warn(args...)
}

func (l *Logrus) Info(mixedArgs ...interface{}) {
	fields, args := l.fields(mixedArgs)
	l.logger.WithFields(fields).Info(args...)
}

func (l *Logrus) Debug(mixedArgs ...interface{}) {
	fields, args := l.fields(mixedArgs)
	l.logger.WithFields(fields).Debug(args...)
}

func (l *Logrus) Trace(mixedArgs ...interface{}) {
	fields, args := l.fields(mixedArgs)
	l.logger.WithFields(fields).Trace(args...)
}

func (l *Logrus) fields(mixedArgs []interface{}) (logrus.Fields, []interface{}) {
	f, args := SplitArgs(mixedArgs)
	return logrus.Fields(f), args
}
