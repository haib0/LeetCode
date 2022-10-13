/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */
package leetcode

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var backtrack func(start, pathsum int, path []int)
	backtrack = func(start, pathsum int, path []int) {
		if len(path) == k || pathsum >= n {
			if len(path) == k && pathsum == n {
				ans = append(ans, append([]int{}, path...))
			}
			return
		}
		for i := start; i < n && i <= 9; i++ {
			backtrack(i+1, pathsum+i, append(path, i))
		}
	}

	backtrack(1, 0, []int{})
	return ans
}

// @lc code=end
