package main

import (
	"flag"
	"fmt"
	"treesource/internal/treesource"
	"treesource/pkg/istty"
)

func main() {
	if !istty.Is() {
		// We close the console window on Windows. This allows us to have a graphical window without a console when running from a non-console location.
		istty.Close()
	}

	treesource.SetupCmds()

	useTUI := flag.Bool("tui", true, "use text interface")
	useGUI := flag.Bool("gui", false, "use graphical interface")

	flag.Usage = func() {
		treesource.ShowHelp()
	}

	flag.Parse()

	if *useGUI || !istty.Is() {
		if err := treesource.RunGUI(); err != nil {
			fmt.Println(err)
		}
	} else if *useTUI || istty.Is() {
		if err := treesource.RunTUI(); err != nil {
			fmt.Println(err)
		}
	}
}
