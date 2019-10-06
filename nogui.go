// +build nogui

package main

import "errors"

func runGUI() error {
	return errors.New("GUI is disabled")
}

func (a *App) Dispatch(e string, v interface{}) {
}
