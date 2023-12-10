package cmd

import (
	"log"

	"github.com/aboxofsox/iggy/pkg/iggy"
	"github.com/spf13/cobra"
)

var all bool
var paths []string

func init() {
	combine.Flags().BoolVarP(&all, "all", "a", false, "Combine all .gitignore files in the current directory")
	combine.Flags().StringSliceVarP(&paths, "paths", "p", []string{}, "Combine .gitignore files at the specified paths")
	rootCmd.AddCommand(combine)
}

var combine = &cobra.Command{
	Use:   "combine",
	Short: "Combine multiple .gitignore files into one",
	Run: func(cmd *cobra.Command, args []string) {
		if all {
			err := iggy.CombineAll()
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		if len(paths) > 0 {
			err := iggy.Combine(paths...)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	},
}
