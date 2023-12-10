package cmd

import (
	"log"

	"github.com/aboxofsox/iggy/pkg/iggy"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(combine)
}

var combine = &cobra.Command{
	Use:   "combine",
	Short: "Combine multiple .gitignore files into one",
	Run: func(cmd *cobra.Command, args []string) {
		err := iggy.CombineAll()
		if err != nil {
			log.Fatal(err)
		}
	},
}
