package importer

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/cristianpontes/teamwork-go-challenge/pkg/csv"
	"github.com/cristianpontes/teamwork-go-challenge/pkg/customer"
	"github.com/cristianpontes/teamwork-go-challenge/pkg/log"
	assertion "github.com/stretchr/testify/assert"
)

func TestNewImporter(t *testing.T) {
	assert := assertion.New(t)

	assert.NotNil(newImporter(nil, nil))
}

func TestImporter_Execute_SimplifiedReport(t *testing.T) {
	assert := assertion.New(t)

	logs := &bytes.Buffer{}

	logger := log.New("importer-integration-test", "cristianpotnes")
	logger.SetOutput(logs)

	cmd := newImporter(logger, customer.NewImporter(csv.NewUnmarshaller()))

	err := cmd.execute("./testing/stubs/customer-import.min.csv", false)
	assert.NoError(err)

	assert.Contains(logs.String(), "\\ncpontes.com -> 1\\n")
	assert.Contains(logs.String(), "\\nhubpages.com -> 2\\n")
}

func TestImporter_Execute_DetailedReport(t *testing.T) {
	assert := assertion.New(t)

	logs := &bytes.Buffer{}

	logger := log.New("importer-integration-test", "cristianpotnes")
	logger.SetOutput(logs)

	cmd := newImporter(logger, customer.NewImporter(csv.NewUnmarshaller()))

	err := cmd.execute("./testing/stubs/customer-import.min.csv", true)
	assert.NoError(err)

	expectedOutput, err := ioutil.ReadFile("./testing/stubs/customer-import.detailed-output.txt")
	assert.NoError(err)
	assert.Contains(strings.TrimSpace(logs.String()), string(expectedOutput))
}
