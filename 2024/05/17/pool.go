package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建一个 sync.Pool 对象
	pool := sync.Pool{
		New: func() interface{} {
			return new(int)
		},
	}

	// 从 pool 中获取对象
	obj := pool.Get().(*int)
	fmt.Println("从 pool 中获取的对象:", *obj)

	// 将对象放回 pool 中
	*obj = 42
	pool.Put(obj)

	// 再次从 pool 中获取对象
	obj2 := pool.Get().(*int)
	fmt.Println("再次从 pool 中获取的对象:", *obj2)
}
