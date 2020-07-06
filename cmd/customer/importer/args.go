package importer

import "github.com/spf13/cobra"

type args struct {
	ImportCSV string
}

func newArgs(cmd *cobra.Command) (*args, error) {
	var a args

	cmd.Flags().StringVarP(
		&a.ImportCSV,
		"file",
		"",
		"",
		"Path to a CSV file containing the list of the customers to be imported.",
	)

	if err := cmd.MarkFlagRequired("file"); err != nil {
		return nil, err
	}

	return &a, nil
}
