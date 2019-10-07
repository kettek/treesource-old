// +build !notui

package main

import (
	"flag"
)

func runTUI() error {
	args := flag.Args()
	if len(args) == 0 {
		showHelp()
		return nil
	}
	cmd, args := args[0], args[1:]

	switch cmd {
	case "init":
		app.Init()
	case "sync":
		app.Sync()
	default:
		showHelp()
	}
	return nil
}
