package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// MinMax 返回给定切片的最小和最大值
type Ordered interface {
	constraints.Ordered
}

// MinMax finds the minimum and maximum values in a slice of Ordered elements.
func MinMax[T Ordered](slice []T) (T, T) {
	if len(slice) == 0 {
		var zero T // The zero value for the type
		return zero, zero
	}
	min, max := slice[0], slice[0]
	for _, value := range slice {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func main() {
	numbers := []int{1, 3, 5, 7, 9, 11}
	min, max := MinMax(numbers)
	fmt.Println("Min:", min, "Max:", max) // 输出: Min: 1 Max: 11

	floats := []float64{3.14, 1.41, 2.71, 1.73}
	fmin, fmax := MinMax(floats)
	fmt.Println("Min:", fmin, "Max:", fmax) // 输出: Min: 1.41 Max: 3.14
}
