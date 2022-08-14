package leetcode

/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设 该网格的四条边均被水包围。
*/
func numDistinctIslands(grid [][]int) int {

	row, col := len(grid), len(grid[0])

	var shape []rune
	var dfs func(x, y int, dir rune)
	dfs = func(x, y int, dir rune) {
		if x < 0 || x >= row || y < 0 || y >= col || grid[x][y] == 0 {
			return
		}

		grid[x][y] = 0
		shape = append(shape, dir)
		dfs(x-1, y, 0)
		dfs(x, y-1, 1)
		dfs(x+1, y, 2)
		dfs(x, y+1, 3)
		shape = append(shape, -dir)
	}

	var ans int
	var islands = make(map[string]bool)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				shape = []rune{}
				dfs(i, j, -1)
				shapeS := string(shape)
				if _, ok := islands[shapeS]; !ok {
					islands[shapeS] = true
					ans++
				}
			}
		}
	}

	return ans
}
