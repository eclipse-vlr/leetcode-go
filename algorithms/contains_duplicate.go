package algorithms

func containsDuplicate(nums []int) bool {
	seen := map[int]bool{}

	for _, value := range nums {
		if !seen[value] {
			seen[value] = true
		} else {
			return true
		}
	}
	return false
}
