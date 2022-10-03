/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */

package leetcode

// @lc code=start

func partitionLabels(s string) []int {
	var lastIndex = make([]int, 26)
	for i, v := range s {
		lastIndex[v-'a'] = i
	}
	var ans []int
	var start int
	for start < len(s) {
		end := start
		for i := start; i < len(s) && i <= end; i++ {
			last := lastIndex[s[i]-'a']
			if end < last {
				end = last
			}
		}
		ans = append(ans, end-start+1)
		start = end + 1
	}
	return ans
}

// @lc code=end
