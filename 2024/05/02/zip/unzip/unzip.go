package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开ZIP文件
	r, err := zip.OpenReader("example.zip")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	// 遍历ZIP包中的每个文件和目录
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(os.Stdout, rc)
		rc.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println()
	}
}
