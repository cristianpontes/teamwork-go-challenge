package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Info(msg string)
	Error(msg error)
	SetOutput(w io.Writer)
}

type Log struct {
	log    *logrus.Logger
	config struct {
		appName string
		vendor  string
	}
}

func New(appName, vendor string) Logger {
	r := &Log{}
	r.log = logrus.New()
	r.config.appName = appName
	r.config.vendor = vendor

	return r
}

func (l *Log) SetOutput(w io.Writer) {
	l.log.Out = w
}

func (l *Log) Info(msg string) {
	l.log.WithFields(l.getFields()).Info(msg)
}

func (l *Log) Error(err error) {
	l.log.WithFields(l.getFields()).Error(err.Error())
}

func (l *Log) getFields() logrus.Fields {
	return logrus.Fields{
		"application_name": l.config.appName,
		"vendor":           l.config.vendor,
	}
}
