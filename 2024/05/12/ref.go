package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := map[string]int{"one": 1, "two": 2}
	b := map[string]int{"one": 1, "two": 2}

	fmt.Println(reflect.DeepEqual(a, b)) // 输出: true

	c := []int{1, 2, 3}
	d := []int{1, 2, 4}

	fmt.Println(reflect.DeepEqual(c, d)) // 输出: false
}
