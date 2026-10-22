package algorithms

import (
	"slices"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "basic case",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			name: "empty string",
			strs: []string{""},
			want: [][]string{{""}},
		},
		{
			name: "single character",
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
		{
			name: "no anagrams",
			strs: []string{"abc", "def"},
			want: [][]string{{"abc"}, {"def"}},
		},
		{
			name: "empty input",
			strs: []string{},
			want: [][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.strs)

			// Сортируем строки внутри каждой группы
			normalize := func(groups [][]string) {
				for _, group := range groups {
					slices.Sort(group)
				}

				// Сортируем сами группы
				slices.SortFunc(groups, func(a, b []string) int {
					return slices.Compare(a, b)
				})
			}

			normalize(got)
			normalize(tt.want)

			if !slices.EqualFunc(got, tt.want, func(a, b []string) bool {
				return slices.Equal(a, b)
			}) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
