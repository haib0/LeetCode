/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */
package leetcode

import (
	"container/heap"
)

// @lc code=start

type MedianFinder struct {
	smaller maxSliceHp
	bigger  minSliceHp
}

func Constructor295() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	if this.smaller.Len() == 0 || num <= this.smaller.IntSlice[0] {
		heap.Push(&this.smaller, num)
		if this.smaller.Len()-1 > this.bigger.Len() {
			x := heap.Pop(&this.smaller)
			heap.Push(&this.bigger, x)
		}
	} else {
		heap.Push(&this.bigger, num)
		if this.smaller.Len() < this.bigger.Len() {
			x := heap.Pop(&this.bigger)
			heap.Push(&this.smaller, x)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.smaller.Len() > this.bigger.Len() {
		return float64(this.smaller.IntSlice[0])
	} else {
		return float64(this.smaller.IntSlice[0]+this.bigger.IntSlice[0]) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end
