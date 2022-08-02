/*
 * @lc app=leetcode.cn id=1436 lang=golang
 *
 * [1436] 旅行终点站
 */
package leetcode

// @lc code=start
func destCity(paths [][]string) string {
	m := make(map[string]bool)
	for _, p := range paths {
		m[p[0]] = true
	}
	for _, p := range paths {
		if !m[p[1]] {
			return p[1]
		}
	}
	return ""
}

// @lc code=end
