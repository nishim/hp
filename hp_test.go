package hp

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestHP(t *testing.T) {
	stdin, _ := os.Open("./testdata/test.html")
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	args := strings.Split("hp", " ")
	expected := []string{
		"meta Description.",
		"title - HP Test Data",
		"HP Test Data.",
		"Extract Pragraphs from HTML.",
		"some image",
		"list element 1",
		"list element 2",
		"list element 3",
		"list element 4",
		"list element 5",
		"table summary",
		"tbody title",
	}

	hp := &HP{InStream: stdin, OutStream: stdout, ErrStream: stderr}
	status := hp.Run(args)
	if status != ExitCodeOK {
		t.Errorf("Exit staus=%d, want %d", status, ExitCodeOK)
	}

	lbl := lineByLine(stdout)
	if len(lbl) != len(expected) {
		t.Errorf("Stdout line count=%d, want %d", len(lbl), len(expected))
	}
	for i := 0; i < len(lbl); i++ {
		if lbl[i] != expected[i] {
			t.Errorf("Line %d is '%s', want '%s", (i + 1), lbl[i], expected[i])
		}
	}
}

func TestHPWithN(t *testing.T) {
	stdin, _ := os.Open("./testdata/test.html")
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	args := strings.Split("hp -n", " ")
	expected := []string{
		"1\tmeta Description.",
		"2\ttitle - HP Test Data",
		"3\tHP Test Data.",
		"4\tExtract Pragraphs from HTML.",
		"5\tsome image",
		"6\tlist element 1",
		"7\tlist element 2",
		"8\tlist element 3",
		"9\tlist element 4",
		"10\tlist element 5",
		"11\ttable summary",
		"12\ttbody title",
	}

	hp := &HP{InStream: stdin, OutStream: stdout, ErrStream: stderr}
	status := hp.Run(args)
	if status != ExitCodeOK {
		t.Errorf("Exit staus=%d, want %d", status, ExitCodeOK)
	}

	lbl := lineByLine(stdout)
	if len(lbl) != len(expected) {
		t.Errorf("Stdout line count=%d, want %d", len(lbl), len(expected))
	}
	for i := 0; i < len(lbl); i++ {
		if lbl[i] != expected[i] {
			t.Errorf("Line %d is '%s', want '%s", (i + 1), lbl[i], expected[i])
		}
	}
}

func lineByLine(r io.Reader) []string {
	var lbl []string
	lbl = make([]string, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lbl = append(lbl, scanner.Text())
	}
	return lbl
}
