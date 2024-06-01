package main

import (
	"reflect"
	"testing"
)

func Test_multiplyMatricesNoCache(t *testing.T) {
	type args struct {
		a [][]int
		b [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiplyMatricesNoCache(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("multiplyMatricesNoCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 基准测试
func BenchmarkMultiplyMatrices(b *testing.B) {
	n := 100
	a := generateRandomMatrix(n)
	bMatrix := generateRandomMatrix(n)
	for i := 0; i < b.N; i++ {
		multiplyMatrices(a, bMatrix)
	}
}

func BenchmarkMultiplyMatricesNoCache(b *testing.B) {
	n := 100
	a := generateRandomMatrix(n)
	bMatrix := generateRandomMatrix(n)
	for i := 0; i < b.N; i++ {
		multiplyMatricesNoCache(a, bMatrix)
	}
}
