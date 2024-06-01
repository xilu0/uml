package main

import "fmt"

func main() {
	// 使用 make 创建切片
	s := make([]string, 0, 5)
	// 使用 append 添加元素
	s = append(s, "Go", "language")
	fmt.Println(s) // 输出: [Go language]

	// 创建复数
	c := complex(5, 7)
	fmt.Println(c) // 输出: (5+7i)

	// 错误处理
	var err error = nil
	if err != nil {
		panic(err)
	}
}
