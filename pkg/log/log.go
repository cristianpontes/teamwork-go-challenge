package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

// Logger represents contract that any logging implementation must follow
type Logger interface {
	Info(msg string)
	Error(msg error)
	SetOutput(w io.Writer)
}

type logger struct {
	log    *logrus.Logger
	config struct {
		appName string
		vendor  string
	}
}

// New creates a new logger using the default implementation
func New(appName, vendor string) Logger {
	r := &logger{}
	r.log = logrus.New()
	r.config.appName = appName
	r.config.vendor = vendor

	return r
}

// SetOutput sets the output to where logs will be written
func (l *logger) SetOutput(w io.Writer) {
	l.log.Out = w
}

// Info logs a message as `info`
func (l *logger) Info(msg string) {
	l.log.WithFields(l.getFields()).Info(msg)
}

// Info logs a message as `error`
func (l *logger) Error(err error) {
	l.log.WithFields(l.getFields()).Error(err.Error())
}

func (l *logger) getFields() logrus.Fields {
	return logrus.Fields{
		"application_name": l.config.appName,
		"vendor":           l.config.vendor,
	}
}
