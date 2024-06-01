package main

import (
	"fmt"
	"runtime"
)

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("Stack:\n%s\n", buf[:n])
}

func main() {
	printStack()
}
