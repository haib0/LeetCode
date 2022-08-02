/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */
package leetcode

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	for m, n := 0, len(matrix[0])-1; m < len(matrix) && n >= 0; {
		if matrix[m][n] == target {
			return true
		} else if matrix[m][n] > target {
			n--
		} else if matrix[m][n] < target {
			m++
		}
	}
	return false
}

// @lc code=end
