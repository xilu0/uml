package main

/*
#include "sum.c"
*/
import "C"
import "fmt"

func main() {
	a := 5
	b := 7
	result := C.sum(C.int(a), C.int(b))
	fmt.Printf("Sum of %d and %d is %d\n", a, b, result)
}
