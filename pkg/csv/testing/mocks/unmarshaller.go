package mocks

import "github.com/stretchr/testify/mock"

type UnmarshallerMock struct {
	mock.Mock
}

func (u *UnmarshallerMock) UnmarshallBytes(in []byte, out interface{}) error {
	args := u.Called(in, out)

	return args.Error(0)
}
