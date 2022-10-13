/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */
package leetcode

// @lc code=start
func partition(s string) [][]string {
	isPalindrome := func(x string) bool {
		i, j := 0, len(x)-1
		for i < j {
			if x[i] != x[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	var ans [][]string
	var backtrack func(start int, trace []string)
	backtrack = func(start int, trace []string) {
		if start == len(s) {
			ans = append(ans, append([]string{}, trace...))
			return
		}
		for end := start + 1; end <= len(s); end++ {
			curr := s[start:end]
			if isPalindrome(curr) {
				backtrack(end, append(trace, curr))
			}
		}
	}

	backtrack(0, []string{})
	return ans
}

// @lc code=end
