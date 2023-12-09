package util

import (
	"os"
	"path/filepath"
)

type HandlerFunc func(string, os.DirEntry)

func ReadDir(dir string, handler HandlerFunc) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		fp := filepath.Join(dir, file.Name())
		if file.IsDir() {
			ReadDir(fp, handler)
		} else {
			handler(fp, file)
		}
	}

	return nil
}
