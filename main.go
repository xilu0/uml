package main

import (
	"fmt"
	"time"
)

func main() {
	var i int
	fmt.Println("Start...")
	for i = 0; i < 5; i++ {
		doSomething(i)
		time.Sleep(time.Second)
	}
	fmt.Println("End...")
}

func doSomething(n int) {
	fmt.Println("Doing something with", n)
}
