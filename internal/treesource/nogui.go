// +build nogui

package treesource

import "errors"

func RunGUI() error {
	return errors.New("GUI is disabled")
}

func (a *App) Dispatch(e string, v interface{}) {
}
