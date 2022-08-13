/*
 * @lc app=leetcode.cn id=384 lang=golang
 *
 * [384] 打乱数组
 */
package leetcode

import (
	"math/rand"
)

// @lc code=start
type Solution384 struct {
	nums, origin []int
}

func Constructor384(nums []int) Solution384 {
	return Solution384{
		nums:   nums,
		origin: append([]int(nil), nums...),
	}
}

func (this *Solution384) Reset() []int {
	copy(this.nums, this.origin)
	return this.nums
}

func (this *Solution384) Shuffle() []int {
	n := len(this.nums)
	for i := 0; i < n; i++ {
		j := i + rand.Intn(n-i)
		this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
	}
	return this.nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end
