/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

package leetcode

// @lc code=start

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	
	m := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	n := len(digits)
	var ans []string

	var backtrace func(i int, path []rune)
	backtrace = func(i int, path []rune) {
		if i == n {
			ans = append(ans, string(path))
			return
		}
		digit := string(digits[i])
		for _, v := range m[digit] {
			backtrace(i+1, append(path, v))
		}

	}

	backtrace(0, []rune{})
	return ans
}

// @lc code=end
