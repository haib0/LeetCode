/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */
package leetcode

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	var ans int
	row, col := len(grid), len(grid[0])

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= row || y < 0 || y >= col || grid[x][y] == 0 {
			return 0
		}
		grid[x][y] = 0
		return dfs(x-1, y) + dfs(x, y-1) + dfs(x+1, y) + dfs(x, y+1) + 1
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				area := dfs(i, j)
				if ans < area {
					ans = area
				}
			}
		}
	}

	return ans
}

// @lc code=end
