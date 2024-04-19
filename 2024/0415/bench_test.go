package benchdemo

import "testing"

// 被测试的函数
func Sum(x, y int) int {
	return x + y
}

// 基准测试函数
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2)
	}
}
