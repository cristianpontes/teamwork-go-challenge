package mail_test

import (
	"testing"

	"github.com/cristianpontes/teamwork-go-challenge/pkg/mail"
	assertion "github.com/stretchr/testify/assert"
)

func TestEmail_Domain(t *testing.T) {
	assert := assertion.New(t)

	assert.Equal("cpontes.com", mail.Email("cristian@cpontes.com").Domain())
	assert.Equal("cpontes.com", mail.Email("cristian+something.else@cpontes.com").Domain())
	assert.Equal("cpontes.com", mail.Email("cristian@something@else@cpontes.com").Domain())
}
