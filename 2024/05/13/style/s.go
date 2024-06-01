package main

import (
	"fmt"
)

// map 函数
func mapFunc(arr []int, f func(int) int) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}

// filter 函数
func filterFunc(arr []int, f func(int) bool) []int {
	result := []int{}
	for _, v := range arr {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	// 将每个元素加1
	mapped := mapFunc(nums, func(n int) int {
		return n + 1
	})
	fmt.Println("Mapped:", mapped) // 输出：[2 3 4 5 6]

	// 过滤出偶数
	filtered := filterFunc(nums, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Filtered:", filtered) // 输出：[2 4]
}
