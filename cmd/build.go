package cmd

import (
	"github.com/aboxofsox/iggy/pkg/iggy"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(build)
}

var build = &cobra.Command{
	Use:   "build",
	Short: "Builds ignore files from .iggy file.",
	Long:  `Builds ignore files from .iggy file.`,
	Run: func(cmd *cobra.Command, args []string) {
		mp, err := iggy.ParseFile(".iggy")
		if err != nil {
			panic(err)
		}

		err = iggy.CreateFiles(mp)
		if err != nil {
			panic(err)
		}
	},
}
