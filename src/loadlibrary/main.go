package main

import (
	"fmt"
	"syscall"
)

func main() {
	loadXXX()
}

func loadXXX() error {
	dllpath := "loadlibrary\\XXX.dll"
	handler, err := syscall.LoadLibrary(dllpath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return err
	}

	defer syscall.FreeLibrary(handler)

	xxx, err := syscall.GetProcAddress(handler, "fnxxx")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return err
	}

	ret, _, _ := syscall.Syscall(xxx, 0, 0, 0, 0)
	if err != nil {
		fmt.Printf("Error: %s \n", err)
	}

	fmt.Println("> fnxxx: ", ret)
	return nil
}
