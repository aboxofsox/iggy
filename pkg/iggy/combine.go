package iggy

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/aboxofsox/iggy/pkg/util"
)

// Combine() combines multiple ignore files.
func Combine(paths ...string) error {
	f, err := os.OpenFile(".iggy", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0700)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, path := range paths {
		key := fmt.Sprintf("@%s{\n", strings.TrimPrefix(filepath.Base(path), "."))
		ff, err := os.OpenFile(path, os.O_RDWR, 0700)
		if err != nil {
			return err
		}

		if _, err := f.WriteString(key); err != nil {
			return err
		}

		scanner := createScanner(ff)
		if err := writeLines(f, scanner); err != nil {
			return err
		}

		if err := writeToFile(f, "}\n\n"); err != nil {
			return err
		}

		ff.Close()

	}

	return nil
}

// CombineAll() combines all ignore files in the current working directory.
func CombineAll() error {
	var paths []string
	util.ReadDir(".", func(path string, file os.DirEntry) {
		if isIgnoreFile(path) {
			paths = append(paths, path)
		}
	})

	return Combine(paths...)
}

func isIgnoreFile(path string) bool {
	return strings.HasPrefix(path, ".") && strings.Contains(path, "ignore")
}
