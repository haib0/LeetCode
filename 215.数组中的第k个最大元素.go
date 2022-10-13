/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */
package leetcode

// @lc code=start
func findKthLargest(nums []int, k int) int {
	partition := func(nums []int, lo, hi int) int {
		pivot := nums[lo]

		i, j := lo+1, hi
		for i <= j {
			for i < hi && nums[i] <= pivot {
				i++
			}
			for j > lo && nums[j] > pivot {
				j--
			}
			if i >= j {
				break
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[lo], nums[j] = nums[j], nums[lo]
		return j
	}

	k = len(nums) - k
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		p := partition(nums, lo, hi)
		if p < k {
			lo = p + 1
		} else if p > k {
			hi = p - 1
		} else {
			return nums[p]
		}
	}

	return -1
}

// @lc code=end
