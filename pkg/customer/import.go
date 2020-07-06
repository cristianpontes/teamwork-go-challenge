package customer

import "github.com/cristianpontes/teamwork-go-challenge/pkg/csv"

type Importer interface {
	FromCSV(csv []byte) (Customers, error)
}

type importer struct {
	csvUnmarshaller csv.Unmarshaller
}

func NewImporter(cu csv.Unmarshaller) Importer {
	return &importer{
		csvUnmarshaller: cu,
	}
}

func (i *importer) FromCSV(csv []byte) (Customers, error) {
	customers := make(Customers, 0)

	err := i.csvUnmarshaller.UnmarshallBytes(csv, &customers)

	return customers, err
}
