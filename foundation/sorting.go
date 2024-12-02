package foundation

import "slices"

func SortIntSlice(unordered []int) []int {
	ordered := make([]int, len(unordered))
	copy(ordered, unordered)
	if slices.IsSorted(ordered) {
		return ordered
	}
	slices.Sort(ordered)
	return ordered
}
