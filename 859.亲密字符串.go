/*
 * @lc app=leetcode.cn id=859 lang=golang
 *
 * [859] 亲密字符串
 */
package leetcode

// @lc code=start
func buddyStrings(s string, goal string) bool {
	if s == goal {
		mapS := make(map[rune]int)
		for _, c := range s {
			if mapS[c] > 0 {
				return true
			}
			mapS[c]++
		}
		return false
	}
	if len(s) != len(goal) {
		return false
	}

	diffS := make([]byte, 2)
	diffG := make([]byte, 2)
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			count++
			if count > 2 {
				return false
			}
			diffS[count-1] = s[i]
			diffG[count-1] = goal[i]
		}
	}

	if diffG[1] == diffS[0] && diffS[1] == diffG[0] {
		return true
	}
	return false
}

// @lc code=end
