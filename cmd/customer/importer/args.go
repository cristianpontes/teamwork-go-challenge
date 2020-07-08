package importer

import "github.com/spf13/cobra"

type args struct {
	ImportCSV      string
	DetailedReport bool
}

func newArgs(cmd *cobra.Command) (*args, error) { // notest
	var a args

	cmd.Flags().StringVarP(
		&a.ImportCSV,
		"file",
		"",
		"",
		"Path to a CSV file containing the list of the customers to be imported.",
	)

	cmd.Flags().BoolVarP(
		&a.DetailedReport,
		"detailed-report",
		"",
		false,
		"Show detailed import report by showing the full list of customers by domain email",
	)

	if err := cmd.MarkFlagRequired("file"); err != nil {
		return nil, err
	}

	return &a, nil
}
