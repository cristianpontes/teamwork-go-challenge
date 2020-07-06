package main

import (
	"os"

	"github.com/cristianpontes/teamwork-go-challenge/cmd/customer/importer"
	"github.com/cristianpontes/teamwork-go-challenge/pkg/csv"
	"github.com/cristianpontes/teamwork-go-challenge/pkg/customer"
	"github.com/cristianpontes/teamwork-go-challenge/pkg/log"
	"github.com/spf13/cobra"
)

func main() {
	lg := log.New("teamwork-go-challenge", "cristianpontes")
	ci := customer.NewImporter(csv.NewUnmarshaller())

	var rootCmd = &cobra.Command{
		Use:   "tw-go-challenge",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Usage(); err != nil {
				lg.Error(err)
				os.Exit(1)
			}
		},
	}

	rootCmd.AddCommand(importer.Command(lg, ci))

	if err := rootCmd.Execute(); err != nil {
		lg.Error(err)
		os.Exit(1)
	}
}
