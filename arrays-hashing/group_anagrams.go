package arrays_hashing

import "slices"

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, word := range strs {
		chars := []byte(word)
		slices.Sort(chars)
		key := string(chars)
		groups[key] = append(groups[key], word)
	}

	result := make([][]string, 0, len(groups))

	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
