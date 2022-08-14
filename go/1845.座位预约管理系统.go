/*
 * @lc app=leetcode.cn id=1845 lang=golang
 *
 * [1845] 座位预约管理系统
 */
package leetcode

import (
	"container/heap"
)

// @lc code=start

// type minSliceHp struct {
// 	sort.IntSlice
// }

// func (h *minSliceHp) Push(x interface{}) {
// 	h.IntSlice = append(h.IntSlice, x.(int))
// }

// func (h *minSliceHp) Pop() interface{} {
// 	n := len(h.IntSlice)
// 	x := h.IntSlice[n-1]
// 	h.IntSlice = h.IntSlice[:n-1]
// 	return x
// }

type SeatManager struct {
	seats minSliceHp
}

func Constructor1845(n int) SeatManager {
	m := SeatManager{}
	for i := 1; i <= n; i++ {
		heap.Push(&m.seats, i)
	}
	return m
}

func (this *SeatManager) Reserve() int {
	return heap.Pop(&this.seats).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(&this.seats, seatNumber)
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
// @lc code=end
