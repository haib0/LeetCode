/*
 * @lc app=leetcode.cn id=304 lang=golang
 *
 * [304] 二维区域和检索 - 矩阵不可变
 */
package leetcode

// @lc code=start
type NumMatrix struct {
	Sums [][]int
}

func Constructor304(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	sums := make([][]int, m+1)
	sums[0] = make([]int, n+1)

	for i, row := range matrix {
		sums[i+1] = make([]int, n+1)
		for j, v := range row {
			sums[i+1][j+1] = v + sums[i][j+1] + sums[i+1][j] - sums[i][j]
		}
	}

	return NumMatrix{
		Sums: sums,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.Sums[row2+1][col2+1] - this.Sums[row2+1][col1] - this.Sums[row1][col2+1] + this.Sums[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end
