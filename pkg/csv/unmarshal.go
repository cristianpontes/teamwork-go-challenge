package csv

import (
	"github.com/gocarina/gocsv"
)

type Unmarshaller interface {
	UnmarshallBytes(in []byte, out interface{}) error
}

type unmarshaller struct{}

func NewUnmarshaller() Unmarshaller {
	return &unmarshaller{}
}

func (*unmarshaller) UnmarshallBytes(in []byte, out interface{}) error {
	return gocsv.UnmarshalBytes(in, out)
}
