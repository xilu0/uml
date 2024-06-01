// main.go

package main

import (
	"fmt"
	"log"

	"github.com/xilu0/uml/2024/06/10/processor"
)

func main() {
	upper := processor.GetProcessor("uppercase")
	if upper == nil {
		log.Fatal("Processor not found: uppercase")
	}
	fmt.Println(upper.Process("hello world"))

	lower := processor.GetProcessor("lowercase")
	if lower == nil {
		log.Fatal("Processor not found: lowercase")
	}
	fmt.Println(lower.Process("HELLO WORLD"))
}
