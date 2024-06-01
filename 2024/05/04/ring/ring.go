package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 创建包含3个元素的环形链表
	r := ring.New(3)

	// 初始化环形链表的值
	n := r.Len()
	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
	// 打印原始链表
	r.Do(func(p interface{}) {
		fmt.Println(p) // 此时会输出0, 1, 2
	})

	// 添加新的环形链表
	s := ring.New(2)
	s.Value = 10
	s = s.Next()
	s.Value = 20
	s.Do(func(p interface{}) {
		fmt.Println(p)
	})
	// 将s链接到r后
	r.Link(s)

	// 打印添加元素后的链表
	r.Do(func(p interface{}) {
		fmt.Println(p)
	})
}
