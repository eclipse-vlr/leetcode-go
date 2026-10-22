package algorithms

import (
	"slices"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"basic", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"single zero", []int{0}, []int{0}},
		{"no zeroes", []int{1, 2, 3}, []int{1, 2, 3}},
		{"all zeroes", []int{0, 0, 0}, []int{0, 0, 0}},
		{"zeroes first", []int{0, 0, 1}, []int{1, 0, 0}},
		{"empty", []int{}, []int{}},
		{"negative values", []int{0, -1, 0, 2}, []int{-1, 2, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.nums)

			if !slices.Equal(tt.nums, tt.want) {
				t.Errorf("got %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
