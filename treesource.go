package main

import (
	"flag"
	"fmt"
)

type App struct {
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

func main() {
	if !isTTY() {
		// We close the console window on Windows. This allows us to have a graphical window without a console when running from a non-console location.
		closeTTY()
	}

	useTUI := flag.Bool("tui", true, "use text interface")
	useGUI := flag.Bool("gui", false, "use graphical interface")
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
