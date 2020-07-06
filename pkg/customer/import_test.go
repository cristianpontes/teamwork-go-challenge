package customer_test

import (
	"errors"
	"testing"

	csvMocks "github.com/cristianpontes/teamwork-go-challenge/pkg/csv/testing/mocks"
	"github.com/cristianpontes/teamwork-go-challenge/pkg/customer"
	"github.com/cristianpontes/teamwork-go-challenge/pkg/mail"
	assertion "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewImporter(t *testing.T) {
	assert := assertion.New(t)

	assert.Implements((*customer.Importer)(nil), customer.NewImporter(nil))
}

func TestImporter_FromCSV_SuccessfulImport(t *testing.T) {
	assert := assertion.New(t)

	mockedCSVContents := []byte("mocked csv contents")

	unmarshaller := &csvMocks.UnmarshallerMock{}
	unmarshaller.
		On("UnmarshallBytes", mockedCSVContents, mock.MatchedBy(func(customers *customer.Customers) bool {
			*customers = append(*customers, &customer.Customer{
				FirstName: "cristian",
				LastName:  "pontes",
				Email:     "cristian@cpontes.com",
				Gender:    "male",
				IPAddress: "127.0.0.1",
			})
			return true
		})).
		Return(nil).
		Once()

	defer unmarshaller.AssertExpectations(t)

	importer := customer.NewImporter(unmarshaller)

	customers, err := importer.FromCSV(mockedCSVContents)
	assert.NoError(err)

	assert.Equal("cristian", customers[0].FirstName)
	assert.Equal("pontes", customers[0].LastName)
	assert.Equal(mail.Email("cristian@cpontes.com"), customers[0].Email)
	assert.Equal("male", customers[0].Gender)
	assert.Equal("127.0.0.1", customers[0].IPAddress)
}

func TestImporter_FromCSV_UnsuccessfulImport(t *testing.T) {
	assert := assertion.New(t)

	mockedCSVContents := []byte("mocked csv contents")
	customers := make(customer.Customers, 0)

	unmarshaller := &csvMocks.UnmarshallerMock{}
	unmarshaller.
		On("UnmarshallBytes", mockedCSVContents, &customers).
		Return(errors.New("something went wrong")).
		Once()

	defer unmarshaller.AssertExpectations(t)

	importer := customer.NewImporter(unmarshaller)

	_, err := importer.FromCSV(mockedCSVContents)
	assert.Error(err)
}
