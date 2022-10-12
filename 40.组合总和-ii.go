/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */
package leetcode

import "sort"

// @lc code=start

func combinationSum2(candidates []int, target int) [][]int {
	var ans [][]int
	sort.Ints(candidates)
	n := len(candidates)
	var visited = make([]bool, n)
	var backtrack func(start int, path []int, sum int)
	backtrack = func(start int, path []int, sum int) {
		if sum > target {
			return
		}
		if sum == target {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := start; i < n; i++ {
			if visited[i] {
				continue
			}
			if i > 0 && candidates[i] == candidates[i-1] && !visited[i-1] {
				continue
			}
			visited[i] = true
			c := candidates[i]
			backtrack(i+1, append(path, c), sum+c) // NOTE: i+1 not start+1
			visited[i] = false
		}
	}

	backtrack(0, []int{}, 0)
	return ans
}

// @lc code=end
