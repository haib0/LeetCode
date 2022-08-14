/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
package leetcode

// @lc code=start
func lengthOfLongestSubstring_map(s string) int {
	m := make(map[byte]int)
	p, r := 0, 0
	for i := 0; i < len(s); i++ {
		value, ok := m[s[i]]
		if ok && i-value < p+1 {
			p = i - value
		} else {
			p = p + 1
		}
		m[s[i]] = i
		if p > r {
			r = p
		}
	}
	return r
}

func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	left, right := 0, 0
	ans := 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		if ans < right-left {
			ans = right - left
		}
	}

	return ans
}

// @lc code=end
