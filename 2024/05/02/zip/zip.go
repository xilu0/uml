package main

import (
	"archive/zip"
	"os"
)

func main() {
	// 创建文件
	outFile, err := os.Create("example.zip")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// 创建zip.Writer
	zw := zip.NewWriter(outFile)
	defer zw.Close()

	// 待添加的文件内容
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This is a readme file."},
		{"hello.txt", "Hello, world!"},
	}

	for _, file := range files {
		f, err := zw.Create(file.Name)
		if err != nil {
			panic(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			panic(err)
		}
	}
}
