package algorithms

import (
	"slices"
	"testing"
)

func TestTwoSumSorted(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{"basic", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"middle", []int{2, 3, 4}, 6, []int{1, 3}},
		{"negative", []int{-1, 0}, -1, []int{1, 2}},
		{"duplicates", []int{1, 1, 3, 5}, 2, []int{1, 2}},
		{"mixed", []int{-5, -2, 0, 3, 9}, 7, []int{2, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSumSorted(tt.numbers, tt.target)

			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
