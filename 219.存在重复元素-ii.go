/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */
package leetcode

// @lc code=start

func containsNearbyDuplicate(nums []int, k int) bool {
	var AbsInt = func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	lastI := make(map[int]int)
	for i, v := range nums {
		x, ok := lastI[v]
		if ok && AbsInt(x-i) <= k {
			return true
		}
		lastI[v] = i
	}
	return false
}

// @lc code=end
