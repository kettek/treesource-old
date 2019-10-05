package main

// #cgo CFLAGS: -g -Wall
// #include "istty.h"
import "C"

func isTTY() bool {
	t := C.isTTY()

	return bool(t)
}
