/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

package leetcode

// @lc code=start
func numIslands(grid [][]byte) int {
	var ans int
	row, col := len(grid), len(grid[0])

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= row || y < 0 || y >= col || grid[x][y] == '0' {
			return
		}
		grid[x][y] = '0'
		dfs(x-1, y)
		dfs(x, y-1)
		dfs(x+1, y)
		dfs(x, y+1)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}

	return ans
}

// @lc code=end
