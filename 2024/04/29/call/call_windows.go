package main

import (
	"fmt"
	"syscall"
)

func main() {
	handle, err := syscall.GetCurrentProcess()
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("handle:", handle)
}
