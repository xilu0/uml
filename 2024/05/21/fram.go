package main

import (
	"fmt"
	"runtime"
	"time"
)

// traceFunc 用于记录函数的执行时间
func traceFunc(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

// recordFunc 调用 runtime.CallersFrames 来记录函数信息
func recordFunc() {
	callers := make([]uintptr, 10)
	n := runtime.Callers(2, callers)
	frames := runtime.CallersFrames(callers[:n])

	for {
		frame, more := frames.Next()
		fmt.Printf("Function: %s\nFile: %s\nLine: %d\n", frame.Function, frame.File, frame.Line)
		if !more {
			break
		}
	}
}

// exampleFunction 是一个示例函数，用于演示性能分析
func exampleFunction() {
	defer traceFunc(time.Now(), "exampleFunction")
	defer recordFunc()

	// 模拟一些工作
	time.Sleep(2 * time.Second)
}

func main() {
	exampleFunction()
}
