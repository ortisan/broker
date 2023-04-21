package log

import (
	"errors"
	"os"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)
}

type logger struct {
	logger *logrus.Logger
}

func (l *logger) Debugf(format string, args ...any) {
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	})

	l.logger.Debugf(format, args)
}

func (l *logger) Infof(format string, args ...any) {
	l.logger.Infof(format, args)
}

func (l *logger) Warnf(format string, args ...any) {
	l.logger.Warnf(format, args)
}

func (l *logger) Errorf(format string, args ...any) {
	l.logger.Errorf(format, args)
}

func (l *logger) Fatalf(format string, args ...any) {
	l.logger.Fatalf(format, args)
}

func (l *logger) Debug(args ...any) {
	l.logger.Debug(args...)
}

func (l *logger) Info(args ...any) {
	l.logger.Info(args...)
}

func (l *logger) Warn(args ...any) {
	l.logger.Warn(args...)
}

func (l *logger) Error(args ...any) {
	l.logger.Error(args...)
}

func (l *logger) Fatal(args ...any) {
	l.logger.Fatal(args...)
}

func NewLogger(logLevel string) (Logger, error) {
	var log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, errors.New("log level is invalid")
	}
	log.SetLevel(level)
	return &logger{logger: log}, nil
}
