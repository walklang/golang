package main

/*
#include <stdio.h>

int add(int a, int b) { return a + b; }

*/
import "C" // 切勿换行输入import

import "fmt"

func main() {
	fmt.Println(C.add(2, 1))
}