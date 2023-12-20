package qk

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestQuicksort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "empty array", args: args{arr: []int{}}, want: []int{}},
		{name: "already sorted array", args: args{arr: []int{1, 2, 3, 4, 5}}, want: []int{1, 2, 3, 4, 5}},
		{name: "reverse sorted array", args: args{arr: []int{5, 4, 3, 2, 1}}, want: []int{1, 2, 3, 4, 5}},
		{name: "random array", args: args{arr: rand.Perm(100)}, want: func() []int {
			arr := rand.Perm(100)
			sort.Ints(arr)
			return arr
		}()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Quicksort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Quicksort() = %v, want %v", got, tt.want)
			}
		})
	}
}
