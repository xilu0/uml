package main

import "fmt"

// 定义一个接口
type Shape interface {
	Area() float64
}

// 定义一个结构体实现接口
type Circle struct {
	Radius float64
}

// 实现接口方法
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	var s Shape
	s = Circle{Radius: 5}
	fmt.Println("Circle Area:", s.Area())
}
