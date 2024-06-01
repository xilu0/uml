package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 创建新链表
	lst := list.New()

	// 在链表尾部添加元素
	lst.PushBack("first")
	lst.PushBack("second")

	// 在链表头部添加元素
	lst.PushFront("third")

	// 遍历链表
	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
