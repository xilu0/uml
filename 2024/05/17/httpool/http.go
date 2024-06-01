package main

import (
	"io"
	"net/http"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		buf := make([]byte, 1024) // 创建一个 1KB 的缓冲区
		return &buf
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	bufPtr := bufferPool.Get().(*[]byte)
	defer bufferPool.Put(bufPtr)
	buf := *bufPtr

	n, _ := io.ReadFull(r.Body, buf)
	w.Write(buf[:n])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
