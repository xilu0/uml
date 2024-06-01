package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	for i := 0; i < 10; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(100))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(n)
	}
}
