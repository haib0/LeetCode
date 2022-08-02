/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 */
package leetcode

// @lc code=start
func countSegments(s string) int {
	n := 0

	last := ' '
	for _, c := range s {
		if c == ' ' && last != ' ' {
			n += 1
		}
		last = c
	}
	if last == ' ' {
		return n
	}

	return n + 1
}

// @lc code=end
