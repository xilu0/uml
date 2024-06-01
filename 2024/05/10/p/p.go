package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[1:4]
	slice2 := arr[2:5]

	fmt.Println(slice1) // 输出: [2 3 4]
	fmt.Println(slice2) // 输出: [3 4 5]

	slice1[1] = 99
	fmt.Println(arr)    // 输出: [1 2 99 4 5]
	fmt.Println(slice2) // 输出: [99 4 5]
}
