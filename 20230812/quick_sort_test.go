// BEGIN: 8f7e2b3c4d5e
package qk

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestQuicksort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "empty array",
			arr:  []int{},
			want: []int{},
		},
		{
			name: "already sorted array",
			arr:  []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "reverse sorted array",
			arr:  []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "random array",
			arr:  rand.Perm(100),
			want: func() []int {
				arr := rand.Perm(100)
				sort.Ints(arr)
				return arr
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Quicksort(tt.arr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Quicksort(%v) = %v, want %v", tt.arr, got, tt.want)
			}
		})
	}
}

func BenchmarkQuicksort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := rand.Perm(1000)
		Quicksort(arr)
	}
}

// END: 8f7e2b3c4d5e
