package leetcode

import "sort"

// min-heap
type minSliceHp struct {
	sort.IntSlice
}

func (h *minSliceHp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *minSliceHp) Pop() interface{} {
	n := len(h.IntSlice)
	x := h.IntSlice[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return x
}

type maxSliceHp struct {
	minSliceHp
}

func (h maxSliceHp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}
