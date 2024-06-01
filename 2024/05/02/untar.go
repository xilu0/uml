package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开TAR文件
	inFile, err := os.Open("example.tar")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	// 创建tar.Reader
	tr := tar.NewReader(inFile)

	// 遍历TAR档案中的所有文件
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // 没有更多的文件了
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			panic(err)
		}
		fmt.Println()
	}
}
