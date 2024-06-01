package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 创建一个用于接收信号的 channel
	sigs := make(chan os.Signal, 1)
	// 注册我们感兴趣的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 创建一个用于通知程序可以结束的 channel
	done := make(chan bool, 1)

	go func() {
		sig := <-sigs // 等待信号
		fmt.Println()
		fmt.Println(sig)
		done <- true // 收到信号，通知程序结束
	}()
	syscall.Exec(os.Args[0], os.Args, os.Environ())
	fmt.Println("awaiting signal")
	<-done // 等待结束通知
	fmt.Println("exiting")
}
