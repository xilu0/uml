package main

import (
	"fmt"
	"strings"
)

func main() {
	// 创建一个新的 Replacer
	replacer := strings.NewReplacer("Hello", "Hi", "World", "Go")

	// 原始字符串
	original := "Hello, World!"

	// 执行替换操作
	result := replacer.Replace(original)

	// 输出结果
	fmt.Println(result) // 输出：Hi, Go!
}
