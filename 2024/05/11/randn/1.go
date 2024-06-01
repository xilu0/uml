package main

import (
	"fmt"
)

const (
	a = 1664525
	c = 1013904223
	m = 1111111111
)

func lcg(seed uint32) func() uint32 {
	state := seed
	return func() uint32 {
		state = (a*state + c) % m
		return state
	}
}

func main() {
	rand := lcg(1)
	for i := 0; i < 10; i++ {
		fmt.Println(rand())
	}
}
