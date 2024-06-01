package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 创建bufio.Reader对象
	reader := bufio.NewReader(file)

	// 使用reader进行读操作
	for {
		// 每次读取一行数据
		line, err := reader.ReadString('\n')
		println(line)
		if err != nil {
			log.Fatal(err)
		}
		// 打印行数据
	}
}
