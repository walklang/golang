package util

/*
#include "util.c"
*/
import "C"

import "fmt"

func GolangSum(a, b int) int {
	s := C.sum(C.int(a), C.int(b))
	fmt.Println(s)
	return a + b 
}