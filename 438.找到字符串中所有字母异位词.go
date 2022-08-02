/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */
package leetcode

// @lc code=start
func findAnagrams(s string, p string) []int {
	var results []int
	if len(s) < len(p) {
		return results
	}
	var sCount, pCount [26]int
	for i, r := range p {
		sCount[s[i]-'a']++
		pCount[r-'a']++
	}
	if sCount == pCount {
		results = append(results, 0)
	}
	for i := 0; i < len(s)-len(p); i++ {
		sCount[s[i]-'a']--
		sCount[s[i+len(p)]-'a']++
		if sCount == pCount {
			results = append(results, i+1)
		}
	}
	return results
}

// @lc code=end
