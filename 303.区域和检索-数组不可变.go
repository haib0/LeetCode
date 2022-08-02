/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */

package leetcode

// @lc code=start
type NumArray struct {
	nums []int
	sums []int
}

func Constructor303(nums []int) NumArray {
	sums := make([]int, len(nums))
	sum := 0
	for i, v := range nums {
		sum += v
		sums[i] = sum
	}
	return NumArray{
		nums: nums,
		sums: sums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.sums[right]
	} else {
		return this.sums[right] - this.sums[left-1]
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end
