package main

import (
	"archive/tar"
	"os"
)

func main() {
	// 创建一个文件用于写入TAR数据
	outFile, err := os.Create("example.tar")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// 创建一个tar.Writer
	tw := tar.NewWriter(outFile)
	defer tw.Close()

	// 添加文件到TAR档案
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}

	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			panic(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			panic(err)
		}
	}
}
