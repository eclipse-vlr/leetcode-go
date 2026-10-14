package arrays_hashing

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "duplicate at end",
			nums: []int{1, 2, 3, 4, 1},
			want: true,
		},
		{
			name: "no duplicates",
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "empty slice",
			nums: []int{},
			want: false,
		},
		{
			name: "duplicate at start",
			nums: []int{1, 1, 2, 3, 4},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsDuplicate(tt.nums)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
