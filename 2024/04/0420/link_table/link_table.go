package main

import "fmt"

// ListNode 定义链表节点
type ListNode struct {
	Value int
	Next  *ListNode
}

// LinkedList 定义链表结构
type LinkedList struct {
	Head *ListNode
}

func (l *LinkedList) InsertAtHead(value int) {
	node := &ListNode{Value: value}
	if l.Head != nil {
		node.Next = l.Head
	}
	l.Head = node
}

func (l *LinkedList) InsertAtTail(value int) {
	node := &ListNode{Value: value}
	if l.Head == nil {
		l.Head = node
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
}

func (l *LinkedList) DeleteHead() {
	if l.Head != nil {
		l.Head = l.Head.Next
	}
}

func (l *LinkedList) Print() {
	current := l.Head
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Println()
}

func main() {
	list := LinkedList{}
	list.InsertAtHead(1)
	list.InsertAtTail(2)
	list.InsertAtHead(0)
	list.Print() // 输出应为 0 1 2

	list.DeleteHead()
	list.Print() // 输出应为 1 2
}
