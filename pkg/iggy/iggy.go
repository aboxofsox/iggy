package iggy

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/aboxofsox/iggy/pkg/util"
)

//go:embed .iggy
var iggyFile string

func CreateIggy() error {
	f, err := os.Create(".iggy")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(iggyFile)
	if err != nil {
		return err
	}

	return nil
}

func ParseFile(path string) (map[string][]string, error) {
	f, err := openFile(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := createScanner(f)

	return parseLines(scanner)
}

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

		if err := writeToFile(f, "\n\n"); err != nil {
			return err
		}

		ff.Close()

		if err := moveFile(path); err != nil {
			return err
		}
	}

	return nil
}

func CombineAll() error {
	var paths []string
	util.ReadDir(".", func(path string, file os.DirEntry) {
		if isIgnoreFile(path) {
			paths = append(paths, path)
		}
	})

	return Combine(paths...)
}

func moveFile(path string) error {
	if _, err := os.Stat("ignore"); os.IsNotExist(err) {
		if err := os.Mkdir("ignore", os.ModeDir); err != nil {
			return err
		}
	}

	fp := filepath.Join("ignore", filepath.Base(path))

	return os.Rename(path, fp)
}

func isIgnoreFile(path string) bool {
	return strings.HasSuffix(path, ".") && strings.Contains(path, "ignore")
}

func writeToFile(f *os.File, s string) error {
	_, err := f.WriteString(s)
	return err
}

func writeLines(f *os.File, scanner *bufio.Scanner) error {
	for scanner.Scan() {
		line := fmt.Sprintf("\t%s\n", scanner.Text())
		if err := writeToFile(f, line); err != nil {
			return err
		}
	}
	return nil
}

func trim(line string) string {
	return strings.TrimSuffix(strings.TrimPrefix(line, "@"), "{")
}

func clean(line string) string {
	return strings.Trim(line, " ")
}

func openFile(path string) (*os.File, error) {
	return os.Open(path)
}

func createScanner(f *os.File) *bufio.Scanner {
	return bufio.NewScanner(f)
}

func isStartOfBlock(line string) bool {
	return strings.HasPrefix(line, "@") && strings.HasSuffix(line, "{")
}

func isEndOfBlock(line string) bool {
	return strings.HasSuffix(line, "}")
}

func addLineToBlock(mp map[string][]string, name string, line string) {
	mp[name] = append(mp[name], clean(line))
}

func parseLines(scanner *bufio.Scanner) (map[string][]string, error) {
	name := ""
	mp := make(map[string][]string)
	startReading := false
	for scanner.Scan() {
		line := scanner.Text()

		if isStartOfBlock(line) {
			startReading = true
			name = trim(line)
			mp[name] = []string{}
			continue
		}

		if startReading {
			addLineToBlock(mp, name, line)
		}

		if isEndOfBlock(line) {
			startReading = false
			continue
		}
	}

	return mp, nil
}
