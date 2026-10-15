package arrays_hashing

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			s:    "aab",
			t:    "abb",
			want: false,
		},
		{
			s:    "abc",
			t:    "ab",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := isAnagram(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
