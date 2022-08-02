/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */
package leetcode

// @lc code=start
func generate(numRows int) [][]int {
	// when numRows=2
	if numRows == 1 {
		return [][]int{{1}}
	}
	// last_result = {{1}}
	last_result := generate(numRows - 1)
	// last_row = {1}
	last_row := last_result[numRows-2]
	this_row := []int{1}
	// numRows-2=0
	for i := 0; i < numRows-2; i++ {
		this_row = append(this_row, last_row[i]+last_row[i+1])
	}
	this_row = append(this_row, 1)
	return append(last_result, this_row)
}

// @lc code=end
