/*
 * @lc app=leetcode.cn id=703 lang=golang
 *
 * [703] 数据流中的第 K 大元素
 */
package leetcode

import (
	"container/heap"
)

// @lc code=start

type KthLargest struct {
	k    int
	topK minSliceHp
}

func Constructor703(k int, nums []int) KthLargest {
	this := KthLargest{k: k}

	for i, v := range nums {
		if i < k {
			heap.Push(&this.topK, v)
		} else {
			this.Add(v)
		}
	}
	return this
}

func (this *KthLargest) Add(val int) int {
	if this.topK.Len() < this.k {
		heap.Push(&this.topK, val)
	} else if val > this.topK.IntSlice[0] {
		heap.Pop(&this.topK)
		heap.Push(&this.topK, val)
	}

	return this.topK.IntSlice[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end
