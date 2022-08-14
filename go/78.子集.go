/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

package leetcode

// @lc code=start

func subsets(nums []int) [][]int {
	n := len(nums)
	var ans [][]int

	var trace []int
	var backtracking func(start int)
	backtracking = func(start int) {
		traceCopy := make([]int, len(trace))
		copy(traceCopy, trace)
		ans = append(ans, traceCopy)
		if len(traceCopy) == n {
			return
		}

		for i := start; i < n; i++ {
			trace = append(trace, nums[i])
			backtracking(i + 1)
			trace = trace[:len(trace)-1]
		}
	}

	backtracking(0)
	return ans
}

// @lc code=end
