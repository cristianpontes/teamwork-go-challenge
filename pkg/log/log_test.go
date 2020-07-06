package log_test

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/cristianpontes/teamwork-go-challenge/pkg/log"

	assertion "github.com/stretchr/testify/assert"
)

func TestLog_Info(t *testing.T) {
	assert := assertion.New(t)
	buf := &bytes.Buffer{}

	l := log.New("teamwork-go-challenge", "cristianpontes")
	l.SetOutput(buf)

	l.Info("test info")
	assert.NotEmpty(buf.String())
	l.SetOutput(os.Stdout)
}

func TestLog_Error(t *testing.T) {
	assert := assertion.New(t)
	buf := &bytes.Buffer{}

	l := log.New("teamwork-go-challenge", "cristianpontes")
	l.SetOutput(buf)

	l.Error(errors.New("test error"))
	assert.NotEmpty(buf.String())
	l.SetOutput(os.Stdout)
}
