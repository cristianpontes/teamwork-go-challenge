package customer_test

import (
	"testing"

	"github.com/cristianpontes/teamwork-go-challenge/pkg/customer"
	assertion "github.com/stretchr/testify/assert"
)

func TestDomainEmailGroupStrategy_Execute(t *testing.T) {
	assert := assertion.New(t)

	strategy := &customer.DomainEmailGroupStrategy{}

	customers := customer.Customers{
		{
			FirstName: "Cristian",
			LastName:  "Pontes",
			Email:     "cristian@cpontes.com",
			Gender:    "male",
			IPAddress: "127.0.0.1",
		},
		{
			Email: "joe@acme.com",
		},
		{
			Email: "james@acme.com",
		},
		{
			Email: "nicole@acme.com",
		},
	}

	grouped := strategy.Execute(customers)

	assert.Len(grouped, 2)
	assert.Len(grouped["cpontes.com"], 1)
	assert.Len(grouped["acme.com"], 3)
	assert.Equal(customers[0], grouped["cpontes.com"][0])
}
