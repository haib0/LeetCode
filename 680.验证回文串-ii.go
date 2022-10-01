/*
 * @lc app=leetcode.cn id=680 lang=golang
 *
 * [680] 验证回文串 II
 */
package leetcode

// @lc code=start
func validPalindrome(s string) bool {
	var isPalindrome = func(s string) bool {
		i, j := 0, len(s)-1
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return isPalindrome(s[i+1:j+1]) || isPalindrome(s[i:j])
		}
		i++
		j--
	}
	return true
}

// @lc code=end
