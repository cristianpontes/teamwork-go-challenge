package mocks

import (
	"github.com/cristianpontes/teamwork-go-challenge/pkg/customer"
	"github.com/stretchr/testify/mock"
)

// ImporterMock - provides a mock for customer.Importer
type ImporterMock struct {
	mock.Mock
}

// FromCSV - provides a mock for customer.Importer.FromCSV
func (i *ImporterMock) FromCSV(csv []byte) (*customer.Customers, error) {
	args := i.Called(csv)

	return args.Get(0).(*customer.Customers), args.Error(1)
}
