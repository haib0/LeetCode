/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 */
package leetcode

import "strings"

// @lc code=start
func reverseWords(s string) string {
	s_l := strings.Split(s, " ")
	for i, j := 0, len(s_l)-1; i < j; i, j = i+1, j-1 {
		s_l[i], s_l[j] = s_l[j], s_l[i]
	}

	result := ""
	for _, r := range s_l {
		if r != "" {
			result += " " + r
		}
	}
	return strings.Trim(result, " ")
}

// @lc code=end
