/*
 * @lc app=leetcode.cn id=417 lang=golang
 *
 * [417] 太平洋大西洋水流问题
 */
package leetcode

// @lc code=start
func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dfs func(i, j int, canReach map[[2]int]bool)
	dfs = func(i, j int, canReach map[[2]int]bool) {
		if canReach[[2]int{i, j}] {
			return
		}
		canReach[[2]int{i, j}] = true
		for _, d := range directions {
			ni, nj := i+d[0], j+d[1]
			if ni < 0 || ni >= m || nj < 0 || nj >= n {
				continue
			}
			if heights[i][j] > heights[ni][nj] {
				continue
			}
			dfs(ni, nj, canReach)
		}
	}

	var canReachP = make(map[[2]int]bool)
	var canReachA = make(map[[2]int]bool)
	for i := 0; i < m; i++ {
		dfs(i, 0, canReachP)
		dfs(i, n-1, canReachA)
	}
	for j := 0; j < n; j++ {
		dfs(0, j, canReachP)
		dfs(m-1, j, canReachA)
	}

	var ans [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if canReachP[[2]int{i, j}] && canReachA[[2]int{i, j}] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

// @lc code=end
