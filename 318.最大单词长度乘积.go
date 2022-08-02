/*
 * @lc app=leetcode.cn id=318 lang=golang
 *
 * [318] 最大单词长度乘积
 */
package leetcode

// @lc code=start
func maxProduct(words []string) int {
	masks := make([]int, len(words))
	for i, word := range words {
		for _, w := range word {
			masks[i] |= 1 << (w - 'a')
		}
	}

	result := 0
	for i, x := range masks {
		for j, y := range masks[:i] {
			if x&y == 0 {
				curr := len(words[i]) * len(words[j])
				if curr > result {
					result = curr
				}
			}
		}
	}

	return result
}

// @lc code=end
