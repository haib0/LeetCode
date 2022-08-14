/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */
package leetcode

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	m := map[int][]rune{}
	i := 0
	flag := 1

	for _, c := range s {
		m[i] = append(m[i], c)
		if i == numRows-1 {
			flag = -1
		}
		if i == 0 {
			flag = 1
		}
		i += flag
	}
	r := ""
	for i := 0; i < numRows; i++ {
		r += string(m[i])
	}
	return r
}

// @lc code=end
