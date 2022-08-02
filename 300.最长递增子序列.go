/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */
package leetcode

// @lc code=start
func lengthOfLIS(nums []int) int {
	l := 0
	piles := make([]int, len(nums))

	for _, n := range nums {
		// binary search n in piles, left bound
		left, right := 0, l
		for left < right {
			mid := (left + right) / 2
			if piles[mid] < n {
				left = mid + 1
			} else {
				right = mid
			}
		}

		if left == l {
			l++
		}

		piles[left] = n
	}
	return l
}

// @lc code=end
