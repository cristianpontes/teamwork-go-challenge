package customer

import "github.com/cristianpontes/teamwork-go-challenge/pkg/csv"

// Importer represents contract that any customer importer implementation must follow
type Importer interface {
	FromCSV(csv []byte) (Customers, error)
}

type importer struct {
	csvUnmarshaller csv.Unmarshaller
}

// NewImporter creates a new customer importer using the default implementation
func NewImporter(cu csv.Unmarshaller) Importer {
	return &importer{
		csvUnmarshaller: cu,
	}
}

// FromCSV takes in the contents of a csv file and unmarshalls it into a slice of customers
func (i *importer) FromCSV(csv []byte) (Customers, error) {
	customers := make(Customers, 0)

	err := i.csvUnmarshaller.UnmarshallBytes(csv, &customers)

	return customers, err
}
