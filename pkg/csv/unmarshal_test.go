package csv_test

import (
	"testing"

	"github.com/cristianpontes/teamwork-go-challenge/pkg/csv"
	assertion "github.com/stretchr/testify/assert"
)

type csvTestStruct struct {
	Foo string `csv:"foo"`
	Bar string `csv:"bar"`
}

func TestNewUnmarshaller(t *testing.T) {
	assert := assertion.New(t)

	assert.Implements((*csv.Unmarshaller)(nil), csv.NewUnmarshaller())
}

func TestUnmarshaller_UnmarshallBytes(t *testing.T) {
	assert := assertion.New(t)

	data := make([]*csvTestStruct, 0)

	unmarshaller := csv.NewUnmarshaller()

	err := unmarshaller.UnmarshallBytes(
		[]byte("foo,bar\nfoo1,bar1\nfoo2,bar2"),
		&data,
	)

	assert.NoError(err)

	assert.Equal("foo1", data[0].Foo)
	assert.Equal("bar1", data[0].Bar)

	assert.Equal("foo2", data[1].Foo)
	assert.Equal("bar2", data[1].Bar)
}
