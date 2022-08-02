/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

package leetcode

// @lc code=start

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	n := len(candidates)
	var backtracking func(start int, com []int, sum int)
	backtracking = func(start int, com []int, sum int) {
		if sum == target {
			// need copy!
			comCopy := make([]int, len(com))
			copy(comCopy, com)
			ans = append(ans, comCopy)
			return
		}
		if sum > target {
			return
		}

		for i := start; i < n; i++ {
			sum = sum + candidates[i]
			com = append(com, candidates[i])
			backtracking(i, com, sum)
			sum -= candidates[i]
			com = com[:len(com)-1]
		}
	}

	backtracking(0, []int{}, 0)

	return ans
}

// @lc code=end
