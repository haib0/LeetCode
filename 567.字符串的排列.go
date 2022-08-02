/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */
package leetcode

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var count1, count2 [26]int
	for i, a := range s1 {
		count1[a-'a']++
		count2[s2[i]-'a']++
	}

	if count1 == count2 {
		return true
	}

	for left := 0; left+len(s1) < len(s2); left++ {
		count2[s2[left]-'a']--
		count2[s2[left+len(s1)]-'a']++
		if count1 == count2 {
			return true
		}
	}

	return false
}

// @lc code=end
