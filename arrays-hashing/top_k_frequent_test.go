package arrays_hashing

import (
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "basic case",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "single element",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "negative numbers",
			nums: []int{-1, -1, 2, 2, 2, 3},
			k:    2,
			want: []int{-1, 2},
		},
		{
			name: "all unique",
			nums: []int{4, 5, 6},
			k:    3,
			want: []int{4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)

			slices.Sort(got)
			slices.Sort(tt.want)

			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
