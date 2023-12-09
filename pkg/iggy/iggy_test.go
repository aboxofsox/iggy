package iggy

import (
	"fmt"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	mp, err := ParseFile(".iggy")
	if err != nil {
		t.Error(err)
	}

	for k, v := range mp {
		f, err := os.OpenFile("."+k, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0700)
		if err != nil {
			t.Error(err)
		}
		for _, v := range v {
			f.WriteString(fmt.Sprintf("%s\n", v))
		}
		f.Close()
	}
}

func TestCombine(t *testing.T) {
	err := Combine("./.gitignore", "./.dockerignore")
	if err != nil {
		t.Error(err)
	}
}
