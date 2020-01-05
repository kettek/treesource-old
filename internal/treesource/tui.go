// +build !notui

package treesource

import (
	"flag"
	"os"
)

func RunTUI() error {
	args := flag.Args()
	if len(args) == 0 {
		ShowHelp()
		return nil
	}
	cmd, args := args[0], args[1:]

	switch cmd {
	case "init":
		if len(args) == 0 {
			dir, err := os.Getwd()
			if err != nil {
				return err
			}
			app.Init(InitCmd{
				TargetDirectory: dir,
			})
		} else {
			app.Init(InitCmd{
				TargetDirectory: args[0],
			})
		}
	case "sync":
		app.Sync(SyncCmd{})
	default:
		ShowHelp()
	}
	return nil
}
