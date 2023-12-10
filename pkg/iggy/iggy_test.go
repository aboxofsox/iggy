package iggy

import (
	"os"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	expect := `# Ignore everything
**

# Except these files/directories
!Dockerfile
!src/

`

	defer func() {
		for _, p := range []string{".gitignore", ".dockerignore", ".eslintignore"} {
			err := os.Remove(p)
			if err != nil {
				t.Fatal(err)
			}
		}
	}()

	mp, err := ParseFile(".iggy")
	if err != nil {
		t.Error(err)
	}

	err = CreateFiles(mp)
	if err != nil {
		t.Error(err)
	}

	f, err := os.ReadFile(".dockerignore")
	if err != nil {
		t.Error(err)
	}

	if strings.EqualFold(string(f), expect) {
		t.Errorf("Expected %s, got %s", expect, string(f))
	}

}

func TestCombine(t *testing.T) {
	err := Combine("./.gitignore", "./.dockerignore", "./.eslintignore")
	if err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(".iggy"); os.IsNotExist(err) {
		t.Error(err)
	}
}

func TestIsIgnore(t *testing.T) {
	ignore := ".gitignore"

	if !isIgnoreFile(ignore) {
		t.Errorf("Expected %s to be an ignore file", ignore)
	}
}
