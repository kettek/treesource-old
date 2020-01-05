package istty

// #cgo CFLAGS: -g -Wall
// #include "istty.h"
import "C"

// Is returns if the process context is a TTY.
func Is() bool {
	t := C.isTTY()

	return bool(t)
}

// Close closes the TTY if possible.
func Close() {
	C.closeTTY()
}
