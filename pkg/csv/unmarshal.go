package csv

import (
	"github.com/gocarina/gocsv"
)

// Unmarshaller represents the contract that any csv unmarshaller must follow
type Unmarshaller interface {
	UnmarshallBytes(in []byte, out interface{}) error
}

type unmarshaller struct{}

// NewUnmarshaller creates a new csv unmarshaller using the default implementation
func NewUnmarshaller() Unmarshaller {
	return &unmarshaller{}
}

// UnmarshallBytes unmarshalls the contents of a csv (in bytes) into a given data structure
func (*unmarshaller) UnmarshallBytes(in []byte, out interface{}) error {
	return gocsv.UnmarshalBytes(in, out)
}
