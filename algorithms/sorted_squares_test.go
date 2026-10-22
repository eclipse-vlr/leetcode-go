package algorithms

import (
	"slices"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "mixed numbers",
			nums: []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "negative numbers",
			nums: []int{-7, -3, -1},
			want: []int{1, 9, 49},
		},
		{
			name: "positive numbers",
			nums: []int{1, 2, 3},
			want: []int{1, 4, 9},
		},
		{
			name: "all zeroes",
			nums: []int{0, 0, 0},
			want: []int{0, 0, 0},
		},
		{
			name: "single element",
			nums: []int{-5},
			want: []int{25},
		},
		{
			name: "empty slice",
			nums: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedSquares(tt.nums)

			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
