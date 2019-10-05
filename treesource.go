package main

import (
	"flag"
	"os"

	"github.com/mattn/go-isatty"
)

type TestState struct {
	Value int `json:"value"`
}

func (t *TestState) Add(n int) {
	t.Value = t.Value + int(n)
}

func (t *TestState) Reset() {
	t.Value = 0
}

func main() {
	useTUI := flag.Bool("tui", false, "use text interface")
	useGUI := flag.Bool("gui", false, "use graphical interface")
	if (*useTUI && !*useGUI) || (isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())) {
		runTUI()
	} else {
		runGUI()
	}
}
