package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type Release struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		BrowserDownloadUrl string `json:"browser_download_url"`
	} `json:"assets"`
}

func init() {
	rootCmd.AddCommand(checkForUpdates)
}

var checkForUpdates = &cobra.Command{
	Use:   "update",
	Short: "Check for updates",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking for updates...")
		err := update()
		if err != nil {
			panic(err)
		}
	},
}

func update() error {
	res, err := http.Get("https://api.github.com/repos/aboxofsox/iggy/releases/latest")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	f, err := os.OpenFile("iggy.exe", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, res.Body)
	if err != nil {
		return err
	}

	return nil
}
