package main

import "sort"

func SortAndTotal(vals []int) (sorted []int, total int) {
	sorted = make([]int, len(vals))

	copy(sorted, vals)
	sort.Ints(sorted)

	for _, val := range sorted {
		total += val
	}

	return
}
