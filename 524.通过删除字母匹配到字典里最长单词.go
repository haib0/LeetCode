/*
 * @lc app=leetcode.cn id=524 lang=golang
 *
 * [524] 通过删除字母匹配到字典里最长单词
 */
package leetcode

import "sort"

// @lc code=start
func findLongestWord(s string, dictionary []string) string {
	var isSub = func(x string) bool {
		i, j := 0, 0
		for i < len(x) && j < len(s) {
			if s[j] == x[i] {
				i++
			}
			j++
		}
		return i == len(x)
	}
	sort.Slice(dictionary, func(i, j int) bool {
		if len(dictionary[i]) == len(dictionary[j]) {
			return dictionary[i] < dictionary[j]
		}
		return len(dictionary[i]) > len(dictionary[j])
	})
	for _, x := range dictionary {
		if isSub(x) {
			return x
		}
	}
	return ""
}

// @lc code=end
