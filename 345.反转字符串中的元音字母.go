/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */

package leetcode

// @lc code=start
func reverseVowels(s string) string {
	var isVowel = func(x byte) bool {
		return x == 'a' || x == 'e' || x == 'i' || x == 'o' || x == 'u' || x == 'A' || x == 'E' || x == 'I' || x == 'O' || x == 'U'
	}
	var ss = []byte(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isVowel(ss[i]) {
			i++
		}
		for i < j && !isVowel(ss[j]) {
			j--
		}
		ss[i], ss[j] = ss[j], ss[i]
		i++
		j--
	}
	return string(ss)
}

// @lc code=end
