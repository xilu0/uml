package main

import (
	"testing"
)

func BenchmarkCompareByInterface(b *testing.B) {
	a := 10
	for i := 0; i < b.N; i++ {
		_, _ = CompareByInterface(a, b)
	}
}

func BenchmarkCompareByGeneric(b *testing.B) {
	a := 10
	for i := 0; i < b.N; i++ {
		_ = CompareByGeneric[int](a, a+1)
	}
}
