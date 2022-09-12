/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
package leetcode

// @lc code=start
func longestPalindrome(s string) string {
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
	n := len(s)
	for len := n; len > 0; len-- {
		for i := 0; i+len <= n; i++ {
			j := i + len
			x := s[i:j]
			if isPalindrome(x) {
				return x
			}
		}
	}
	return s[:1]
}

func longestPalindromedp(s string) string {
	var n = len(s)
	var dp = make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	var maxLen, start = 1, 0
	for len := 2; len <= n; len++ {
		for i := 0; i < n-len+1; i++ {
			j := i + len - 1
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if len <= 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && maxLen < len {
				maxLen = len
				start = i
			}
		}
	}
	return s[start : start+maxLen]
}

// @lc code=end
