package importer

import (
	"os"

	"github.com/cristianpontes/teamwork-go-challenge/pkg/customer"
	"github.com/cristianpontes/teamwork-go-challenge/pkg/log"
	"github.com/spf13/cobra"
)

func Command(
	lg log.Logger,
	ci customer.Importer,
) *cobra.Command {
	c := cobra.Command{
		Use:   "importer",
		Short: "Customer importer reads from a csv file and returns a sorted list of email domains along with the number of customers with e-mail addresses for each domain.",
	}

	args, err := newArgs(&c)
	if err != nil {
		panic(err)
	}

	c.Run = func(*cobra.Command, []string) {
		importer := NewImporter(lg, ci)

		if err := importer.Execute(args.ImportCSV, false); err != nil {
			lg.Error(err)
			os.Exit(1)
		}
	}

	return &c
}
