package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2) // 设置使用的最大CPU数量
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine 1:", i)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine 2:", i)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(6 * time.Second) // 等待goroutine执行完毕
}
