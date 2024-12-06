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

func Contains[S ~[]E, E comparable](listly S, value E) (bool, int) {
	for i, item := range listly {
		if value == item {
			return true, i
		}
	}
	return false, -1
}

func ContainsItemInList[S ~[]E, E comparable](listly S, values S) (bool, int, E) {
	for _, item := range values {
		check, pos := Contains(listly, item)
		if !check {
			continue
		}
		return check, pos, item
	}
	return false, -1, values[0]
}

func Count[S ~[]E, E comparable](listly S, value E) int {
	count := 0
	for _, item := range listly {
		if value == item {
			count += 1
		}
	}
	return count
}
