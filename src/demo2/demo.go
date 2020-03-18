package main

/*
#cgo CFLAGS: -I ./
#cgo LDFLAGS: -L . -ldemo
#include "demo.h"
*/
import "C"

import "fmt"

func main() {
	cmd := C.CString("message")
	C.xprint(&cmd)
}