package main

import (
	"fmt"
	"syscall"
)

func main() {
	// 准备数据
	s := "Hello, Windows!\n"
	bytes := []byte(s)

	// 获取标准输出的句柄
	stdout := syscall.Stdout

	// 使用 syscall 包调用 WriteFile 函数
	var written uint32
	err := syscall.WriteFile(stdout, bytes, &written, nil)
	if err != nil {
		fmt.Printf("WriteFile failed: %v\n", err)
		return
	}

	fmt.Println("Data written successfully.")
}
