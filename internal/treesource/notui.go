// +build notui

package treesource

import "errors"

func RunTUI() error {
	return errors.New("TUI is disabled")
}
