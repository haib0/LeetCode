/*
 * @lc app=leetcode.cn id=451 lang=golang
 *
 * [451] 根据字符出现频率排序
 */
package leetcode

// @lc code=start
func frequencySort(s string) string {
	var cToFreq = make(map[rune]int)
	for _, v := range s {
		cToFreq[v]++
	}
	var freqToCs = make(map[int][]rune)
	for k, v := range cToFreq {
		freqToCs[v] = append(freqToCs[v], k)
	}
	var ans []rune
	for freq := len(s); freq > 0; freq-- {
		v, ok := freqToCs[freq]
		if ok {
			for _, c := range v {
				for i := 0; i < freq; i++ {
					ans = append(ans, c)
				}
			}
		}
	}
	return string(ans)
}

// @lc code=end
