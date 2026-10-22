package algorithms

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		want   []int
		target int
	}{
		{
			name:   "basic case",
			nums:   []int{2, 7, 11, 15},
			want:   []int{0, 1},
			target: 9,
		},
		{
			name:   "middle elements",
			nums:   []int{3, 2, 4},
			want:   []int{1, 2},
			target: 6,
		},
		{
			name:   "same values",
			nums:   []int{3, 3},
			want:   []int{0, 1},
			target: 6,
		},
		{
			name:   "negative numbers",
			nums:   []int{-3, 4, 3, 90},
			want:   []int{0, 2},
			target: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			slices.Sort(got)
			w := slices.Clone(tt.want)
			slices.Sort(w)
			if !slices.Equal(got, w) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
