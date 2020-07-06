package mail

import (
	"strings"
)

type Email string

func (e Email) Domain() string {
	eStr := string(e)

	i := strings.LastIndexByte(eStr, '@')

	return eStr[i+1:]
}
