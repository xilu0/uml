// lib/lib.go
package lib

import "fmt"

var (
	// 包级变量初始化
	LibraryName = "Go Library"
)

func init() {
	// init函数
	fmt.Println("lib包初始化: ", LibraryName)
}

func PrintLibName() {
	fmt.Println("库名: ", LibraryName)
}
