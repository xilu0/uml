package main

import (
	"fmt"
	"syscall"
)

func main() {
	pid, _, err := syscall.Syscall(syscall.SYS_GETPID, 0, 0, 0)
	if err != 0 {
		fmt.Println("系统调用出错:", err)
	} else {
		fmt.Println("当前进程 PID:", pid)
	}
}
