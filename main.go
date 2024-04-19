// main.go
package main

import (
	"fmt"

	"github.com/xilu0/uml/lib" // 导入自定义的lib包
)

var (
	// 包级变量初始化
	appName = "Go Application"
)

func init() {
	// init函数
	fmt.Println("main包初始化: ", appName)
}

func main() {
	fmt.Println("main函数执行")
	lib.PrintLibName()
}
