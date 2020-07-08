package mail

import (
	"strings"
)

// Email - core domain structure
type Email string

// Domain returns the domain address of a given mail.Email
func (e Email) Domain() string {
	eStr := string(e)

	i := strings.LastIndexByte(eStr, '@')

	return eStr[i+1:]
}
