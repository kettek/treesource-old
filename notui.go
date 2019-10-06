// +build notui

package main

import "errors"

func runTUI() error {
	return errors.New("TUI is disabled")
}
