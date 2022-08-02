/*
 * @lc app=leetcode.cn id=1905 lang=golang
 *
 * [1905] 统计子岛屿
 */

package leetcode

// @lc code=start
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid2), len(grid2[0])
	var island [][]int
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid2[i][j] == 0 {
			return
		}

		island = append(island, []int{i, j})
		grid2[i][j] = 0

		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				island = [][]int{}
				dfs(i, j)
				ok := true
				for _, v := range island {
					if grid1[v[0]][v[1]] == 0 {
						ok = false
						break
					}
				}
				if ok {
					ans++
				}
			}
		}
	}
	return ans
}

// @lc code=end
