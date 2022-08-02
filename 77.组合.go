/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

package leetcode

// @lc code=start

func combine(n int, k int) [][]int {
	var ans [][]int

	var trace []int
	var backtracking func(start int)
	backtracking = func(start int) {
		if len(trace) == k {
			traceCopy := make([]int, len(trace))
			copy(traceCopy, trace)
			ans = append(ans, traceCopy)
			return
		}

		for i := start; i <= n; i++ {
			trace = append(trace, i)
			backtracking(i + 1)
			trace = trace[:len(trace)-1]
		}
	}

	backtracking(1)
	return ans
}

// @lc code=end
