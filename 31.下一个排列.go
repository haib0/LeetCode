/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */
package leetcode

import "sort"

// @lc code=start

// Like [4, 5, 6, 2, 1, 7, 3],
// find the first decreate number from right: 1
// change with the smallest bigger number right to it: 3
// ---> [4, 5, 6, 2, 3, 7, 1],
// then sort numbers to right
// ---> [4, 5, 6, 2, 3, 1, 7]
func nextPermutation(nums []int) {
	n := len(nums)
	var i int
	for i = n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i == -1 {
		sort.Ints(nums)
		return
	}
	j := i + 1
	for ii := i + 2; ii < n; ii++ {
		if nums[i] < nums[ii] && nums[ii] < nums[j] {
			j = ii
		}
	}
	nums[i], nums[j] = nums[j], nums[i]
	sort.Ints(nums[i+1:])
}

// @lc code=end
