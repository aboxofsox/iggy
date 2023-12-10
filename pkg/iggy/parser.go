package iggy

import (
	"bufio"
)

// ParseFile() parses the .iggy file and returns a map of file names and their respecdtive lines.
func ParseFile(path string) (map[string][]string, error) {
	f, err := openFile(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := createScanner(f)

	return parseLines(scanner)
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
