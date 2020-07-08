package mocks

import "github.com/stretchr/testify/mock"

// UnmarshallerMock - provides a mock for csv.Unmarshaller
type UnmarshallerMock struct {
	mock.Mock
}

// UnmarshallBytes - provides a mock for csv.Unmarshaller.UnmarshallBytes
func (u *UnmarshallerMock) UnmarshallBytes(in []byte, out interface{}) error {
	args := u.Called(in, out)

	return args.Error(0)
}
