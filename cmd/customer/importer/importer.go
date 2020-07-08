package importer

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cristianpontes/teamwork-go-challenge/pkg/customer"
	"github.com/cristianpontes/teamwork-go-challenge/pkg/log"
)

type importer struct {
	logger   log.Logger
	importer customer.Importer
}

func newImporter(lg log.Logger, ci customer.Importer) *importer {
	return &importer{
		logger:   lg,
		importer: ci,
	}
}

func (i *importer) execute(filePath string, detailedReport bool) error {
	absPath, err := filepath.Abs(filePath)
	if err != nil { // notest
		return err
	}

	csvContents, err := ioutil.ReadFile(filepath.Clean(absPath))
	if err != nil { // notest
		return err
	}

	customers, err := i.importer.FromCSV(csvContents)
	if err != nil { // notest
		return err
	}

	gs := &customer.DomainEmailGroupStrategy{}

	groups := gs.Execute(customers)

	if detailedReport {
		return i.generateDetailedReport(groups)
	}

	return i.generateSimplifiedReport(groups)
}

func (i *importer) generateSimplifiedReport(groups customer.GroupedCustomers) error {
	var report strings.Builder

	report.WriteString("\n")

	for domain, group := range groups {
		report.WriteString(domain)
		report.WriteString(" -> ")
		report.WriteString(strconv.Itoa(len(group)))
		report.WriteString("\n")
	}

	report.WriteString("\n")

	i.logger.Info(report.String())

	return nil
}

func (i *importer) generateDetailedReport(groups customer.GroupedCustomers) error {
	report, err := json.MarshalIndent(groups, "", "  ")
	if err != nil {
		return err
	}

	i.logger.Info(string(report))

	return nil
}
