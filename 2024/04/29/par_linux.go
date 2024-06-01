package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	s := "hello"
	// 获取字符串首地址
	addr := unsafe.Pointer(&s)
	// 转换为 uintptr
	ptr := uintptr(addr)

	// 系统调用，例如某种内存操作，这里仅为示例
	syscall.Syscall(syscall.SYS_WRITE, uintptr(1), ptr, uintptr(len(s)))
	fmt.Println("Done")
}
