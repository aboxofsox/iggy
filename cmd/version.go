package cmd

import "github.com/spf13/cobra"

const Version = "1.0.0"

func init() {
	rootCmd.AddCommand(version)
}

var version = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(Version)
	},
}
