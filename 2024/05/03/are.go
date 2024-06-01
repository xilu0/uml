package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// 打开文件用于写入
	file, err := os.Create("example_output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 创建bufio.Writer对象
	writer := bufio.NewWriter(file)

	// 使用writer写入数据
	_, err = writer.WriteString("Hello, World!")
	if err != nil {
		log.Fatal(err)
	}

	// 刷新缓冲区，确保所有数据都写入底层io.Writer
	writer.Flush()
}
