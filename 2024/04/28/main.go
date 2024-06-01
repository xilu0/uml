package main

func AddWithInterface(items []interface{}) interface{} {
	var sum int
	for _, item := range items {
		num, ok := item.(int)
		if !ok {
			// 运行时错误处理
		}
		sum += num
	}
	return sum
}

// 使用泛型
func AddWithGenerics[T int | float64](items []T) T {
	var sum T
	for _, item := range items {
		sum += item
	}
	return sum
}

func main() {
	var nums = []int{1, 2, 3}
	var items = []interface{}{1, 2.0}
	var _ int = AddWithGenerics(nums)           // 正确
	var _ interface{} = AddWithInterface(items) // 正确
	// var strs = []string{"hello", "world"}
	// var _ int = AddWithGenerics(strs)           // 编译错误：类型不匹配
	var _ interface{} = AddWithInterface(items) // 编译错误：类型不匹配
}
