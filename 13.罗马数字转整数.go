/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */
package leetcode

// @lc code=start
func romanToInt(s string) int {
	romanMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	n := len(s)
	ans := romanMap[s[n-1]]
	for i := n - 2; i >= 0; i-- {
		if romanMap[s[i]] < romanMap[s[i+1]] {
			ans -= romanMap[s[i]]
		} else {
			ans += romanMap[s[i]]
		}
	}
	return ans
}

// @lc code=end
