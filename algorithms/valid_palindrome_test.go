package algorithms

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "classic palindrome",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "not palindrome",
			s:    "race a car",
			want: false,
		},
		{
			name: "empty string",
			s:    "",
			want: true,
		},
		{
			name: "only punctuation",
			s:    ".,!?",
			want: true,
		},
		{
			name: "mixed case",
			s:    "Aa",
			want: true,
		},
		{
			name: "digits",
			s:    "12321",
			want: true,
		},
		{
			name: "different digits",
			s:    "12345",
			want: false,
		},
		{
			name: "unicode",
			s:    "А роза упала на лапу Азора",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.s)

			if got != tt.want {
				t.Errorf("isPalindrome(%q) = %v, want %v",
					tt.s, got, tt.want)
			}
		})
	}
}
