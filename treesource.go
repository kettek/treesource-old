package main

import (
	"flag"
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
	useTUI := flag.Bool("tui", true, "use text interface")
	useGUI := flag.Bool("gui", false, "use graphical interface")
	flag.Parse()

	if *useGUI || !isTTY() {
		runGUI()
	} else if *useTUI || isTTY() {
		runTUI()
	}
}
