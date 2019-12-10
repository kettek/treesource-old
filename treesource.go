package main

import (
	"flag"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"os"
	"path/filepath"
)

type App struct {
	gui     bool
	Title   string
	Entries []AppEntry
}

func (a *App) Init(cmd InitCmd) (err error) {
	dir, err := filepath.Abs(cmd.TargetDirectory)
	if err != nil {
		return err
	}
	fmt.Printf("initializing treesource in \"%s\"\n", dir)
	if cmd.CommandIndex > 0 {
		a.Dispatch("init", cmd)
	}
	return nil
}

func (a *App) Sync(cmd SyncCmd) {
	fmt.Println("sync called")
	a.Dispatch("sync", cmd)
}

func (a *App) Search(cmd SearchCmd) {
	fmt.Println("search called")
	fmt.Printf("searching for %s\n", cmd.SearchString)
	// a.Dispatch("searchResults", ...)
}

type InitCmd struct {
	TargetDirectory string
	CommandIndex int
}

type SyncCmd struct {
	CommandIndex int
}

type SearchCmd struct {
	SearchString string
	CommandIndex int
}

func (a *App) HandleEvent(s string, v interface{}) {
	switch s {
	case "sync":
		var syncCmd SyncCmd
		if err := mapstructure.Decode(v, &syncCmd); err == nil {
			a.Sync(syncCmd)
		}
	case "init":
		var initCmd InitCmd
		if err := mapstructure.Decode(v, &initCmd); err == nil {
			a.Init(initCmd)
		}
	case "search":
		var searchCmd SearchCmd
		if err := mapstructure.Decode(v, &searchCmd); err == nil {
			a.Search(searchCmd)
		}
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

var cmdsHelp map[string]string

func setupCmds() {
	cmdsHelp = make(map[string]string)
	cmdsHelp["init"] = "Initialize the current directory or provided path as a treesource directory"
	cmdsHelp["sync"] = "Synchronize the symbolic link structure with the treesource database"
}

func showHelp() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	for k, v := range cmdsHelp {
		fmt.Fprintf(flag.CommandLine.Output(), "  %s\n\t%s\n", k, v)
	}
}

func main() {
	if !isTTY() {
		// We close the console window on Windows. This allows us to have a graphical window without a console when running from a non-console location.
		closeTTY()
	}

	setupCmds()

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
