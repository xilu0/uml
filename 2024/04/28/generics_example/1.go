package main

import "fmt"

// Filter 是一个泛型函数，可以处理任何类型的切片
func Filter[T any](slice []T, fn func(T) bool) []T {
	var result []T
	for _, value := range slice {
		if fn(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	// 整数过滤
	ints := []int{1, 2, 3, 4, 5, 6}
	even := Filter(ints, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(even) // 输出: [2 4 6]

	// 字符串过滤
	strings := []string{"apple", "banana", "grape", "plum"}
	hasP := Filter(strings, func(s string) bool {
		return s[0] == 'p'
	})
	fmt.Println(hasP) // 输出: [plum]
}
