package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

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
			log.Fatal(err.Error())
		}
	},
}

func update() error {
	var releases []*Release

	res, err := http.Get("https://api.github.com/repos/aboxofsox/iggy/releases")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = decodeJSON(res.Body, &releases)
	if err != nil {
		return err
	}

	fmt.Printf("Latest version: %s\n", releases[0].TagName)

	releaseUrl, err := findSystemRelease("1.0.0", releases) // TODO: get version from somewhere or latest
	if err != nil {
		return err
	}

	if _, err := os.Stat(installDir()); os.IsNotExist(err) {
		if err := os.Mkdir(installDir(), os.ModeDir); err != nil {
			return err
		}
	}

	name := "iggy"
	if runtime.GOOS == "windows" {
		name = name + ".exe"
	}
	err = downloadFile(releaseUrl, filepath.Join(installDir(), name))
	if err != nil {
		return err
	}

	return nil
}

func decodeJSON(r io.Reader, v any) error {
	return json.NewDecoder(r).Decode(v)
}

func downloadFile(url, filename string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%s - %d", res.Status, res.StatusCode)
	}

	f, err := os.Create(filename)
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

// func moveFile(from, to string) error {
// 	return os.Rename(from, to)
// }

func installDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return os.TempDir()
	}
	return filepath.Join(home, "iggy")
}

func findSystemRelease(tagname string, rs []*Release) (string, error) {
	for _, r := range rs {
		if !strings.EqualFold(r.TagName, tagname) {
			continue
		}

		if len(r.Assets) == 0 {
			continue
		}
		for _, asset := range r.Assets {
			if strings.Contains(asset.BrowserDownloadUrl, runtime.GOOS) {
				return asset.BrowserDownloadUrl, nil
			}
		}
	}

	return "", fmt.Errorf("no release for %s", runtime.GOOS)
}

// func base(p string) string {
// 	u, err := url.Parse(p)
// 	if err != nil {
// 		return ""
// 	}
// 	return path.Base(u.Path)
// }
