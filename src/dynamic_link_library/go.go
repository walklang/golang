package main

import "C"

//export Sum
func Sum(a int, b int) int { //最简单的计算和
	return a + b
}

// Need a main function to make CGO compile package as C shared library
func main() {
}
