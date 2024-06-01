package main

import (
	"fmt"
)

// 定义一个函数类型
type operation func(int, int) int

// 高阶函数，接收一个函数作为参数
func compute(a int, b int, op operation) int {
	return op(a, b)
}

// 加法操作
func add(a int, b int) int {
	return a + b
}

// 主函数
func main() {
	result := compute(3, 4, add)
	fmt.Println("Result:", result) // 输出：Result: 7
}
