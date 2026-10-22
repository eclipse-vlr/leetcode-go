package algorithms

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for indx, value := range nums {
		find := target - value
		index, exists := seen[find]
		if !exists {
			seen[value] = indx
		} else {
			return []int{indx, index}
		}
	}
	return nil
}
