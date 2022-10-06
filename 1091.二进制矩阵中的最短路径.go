/*
 * @lc app=leetcode.cn id=1091 lang=golang
 *
 * [1091] 二进制矩阵中的最短路径
 */
package leetcode

// @lc code=start
func shortestPathBinaryMatrix(grid [][]int) int {
	var ans int
	var steps = [][]int{{1, -1}, {1, 0}, {1, 1}, {0, -1}, {0, 1}, {-1, -1}, {-1, 0}, {-1, 1}}
	var queue = [][]int{{0, 0}}
	for len(queue) > 0 {
		ans++
		n := len(queue)
		for i := 0; i < n; i++ {
			x := queue[i][0]
			y := queue[i][1]
			if grid[x][y] == 0 {
				if x == len(grid)-1 && y == len(grid[0])-1 {
					return ans
				}
				grid[x][y] = 1
				for _, step := range steps {
					xn := x + step[0]
					yn := y + step[1]
					if 0 <= xn && xn < len(grid) && 0 <= yn && yn < len(grid[0]) {
						queue = append(queue, []int{xn, yn})
					}
				}
			}
		}
		queue = queue[n:]
	}
	return -1
}

// @lc code=end
