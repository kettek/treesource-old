package main

import (
	"flag"
	"fmt"
	"os"
)

type App struct {
	gui     bool
	Title   string
	Entries []AppEntry
}

func (a *App) Init() {
	fmt.Println("init called")
	a.Dispatch("init", nil)
}

func (a *App) Sync() {
	fmt.Println("sync called")
	a.Dispatch("sync", nil)
}

func (a *App) HandleEvent(s string, v interface{}) {
	fmt.Printf("e: %s\n", s)
	switch s {
	case "sync":
		a.Sync()
	case "init":
		a.Init()
	}
}

type AppEntry struct {
	Filename string
	Tags     []string
	Checksum uint32 // CRC32 checksum
}

type AppDatabase struct {
	Tags map[string]AppEntry
}

var app App

func showHelp() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	if !isTTY() {
		// We close the console window on Windows. This allows us to have a graphical window without a console when running from a non-console location.
		closeTTY()
	}

	useTUI := flag.Bool("tui", true, "use text interface")
	useGUI := flag.Bool("gui", false, "use graphical interface")

	flag.Usage = func() {
		showHelp()
	}

	flag.Parse()

	if *useGUI || !isTTY() {
		if err := runGUI(); err != nil {
			fmt.Println(err)
		}
	} else if *useTUI || isTTY() {
		if err := runTUI(); err != nil {
			fmt.Println(err)
		}
	}
}
