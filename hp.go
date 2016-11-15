package hp

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"regexp"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const (
	ExitCodeOK = iota
	ExitCodeError
)

type HP struct {
	InStream             io.Reader
	OutStream, ErrStream io.Writer
	i                    int
	n                    bool
}

func (hp *HP) Run(args []string) int {
	// Initialize.
	hp.i = 1

	// Parse arguments.
	flags := flag.NewFlagSet("hp", flag.ContinueOnError)
	flags.SetOutput(hp.ErrStream)
	flags.BoolVar(&hp.n, "n", false, "Number the output lines, starting at 1.")
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Get html text and parse.
	var htmltext string
	scanner := bufio.NewScanner(hp.InStream)
	for scanner.Scan() {
		htmltext = htmltext + removeDecoElements(scanner.Text())
	}
	doc, err := html.Parse(strings.NewReader(htmltext))
	if err != nil {
		hp.err(err)
		return ExitCodeError
	}

	hp.traverseAndOut(doc)

	return ExitCodeOK
}

func (hp *HP) traverseAndOut(node *html.Node) {
	if node.DataAtom == atom.Script ||
		node.DataAtom == atom.Noscript {
		return
	}

	if node.Type == html.TextNode && strings.TrimSpace(node.Data) != "" {
		hp.out(strings.TrimSpace(node.Data))
	}
	for _, attr := range node.Attr {
		switch attr.Key {
		case atom.Alt.String(), atom.Summary.String(), atom.Title.String():
			hp.out(fmt.Sprintf("%s", attr.Val))
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		hp.traverseAndOut(child)
	}
}

func (hp *HP) out(s string) {
	if hp.n {
		fmt.Fprintln(hp.OutStream, fmt.Sprintf("%d\t%s", hp.i, s))
	} else {
		fmt.Fprintln(hp.OutStream, s)
	}
	hp.i++
}

func (hp *HP) err(a ...interface{}) {
	fmt.Fprintln(hp.ErrStream, a)
}

func removeDecoElements(s string) string {
	r := regexp.MustCompile(`</?(a|span|strong|br|hr)(| [^>]*)>`)
	return r.ReplaceAllString(s, "")
}
