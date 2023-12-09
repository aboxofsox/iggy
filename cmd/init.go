package cmd

import (
	"log"

	"github.com/aboxofsox/iggy/pkg/iggy"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new .iggy file",
	Run: func(cmd *cobra.Command, args []string) {
		err := iggy.CreateIggy()
		if err != nil {
			log.Fatal(err)
		}
	},
}
