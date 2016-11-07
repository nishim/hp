package main

import (
	"os"

	"github.com/nishim/hp"
)

func main() {
	hp := &hp.HP{InStream: os.Stdin, OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(hp.Run(os.Args))
}
