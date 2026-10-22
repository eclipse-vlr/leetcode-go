package algorithms

import "slices"

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++
	}

	type item struct {
		num   int
		count int
	}

	items := make([]item, 0, len(counts))

	for num, count := range counts {
		items = append(items, item{
			num:   num,
			count: count,
		})
	}

	slices.SortFunc(items, func(a, b item) int {
		return b.count - a.count
	})

	result := make([]int, 0, k)

	for i := 0; i < k; i++ {
		result = append(result, items[i].num)
	}

	return result
}
