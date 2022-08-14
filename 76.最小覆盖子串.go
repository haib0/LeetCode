/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */
package leetcode

// @lc code=start
func minWindow(s string, t string) string {
	need, window := make(map[rune]int), make(map[rune]int)
	for _, c := range t {
		need[c]++
	}

	left := 0
	valid := 0
	start, end := 0, len(s)
	for right, c := range s {
		window[c]++
		if window[c] == need[c] {
			valid++
		}
		for valid == len(need) {
			if l := right - left + 1; l < end-start+1 {
				start = left
				end = right
			}
			d := rune(s[left])
			left++
			if window[d] == need[d] {
				valid--
			}
			window[d]--
		}
	}
	if end == len(s) {
		return ""
	}
	return s[start : end+1]
}

// @lc code=end
