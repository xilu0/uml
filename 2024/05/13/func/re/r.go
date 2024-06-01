package main

import (
	"fmt"
)

// 返回一个函数的高阶函数
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// 主函数
func main() {
	multiplier := createMultiplier(2)
	fmt.Println(multiplier(3)) // 输出：6
}
