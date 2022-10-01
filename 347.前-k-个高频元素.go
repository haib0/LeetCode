/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */
package leetcode

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	var numToFreq = make(map[int]int)
	for _, v := range nums {
		numToFreq[v]++
	}
	var freqToNums = make(map[int][]int)
	for num, freq := range numToFreq {
		freqToNums[freq] = append(freqToNums[freq], num)
	}
	var ans []int
	for freq := len(nums); freq > 0; freq-- {
		v, ok := freqToNums[freq]
		if ok {
			ans = append(ans, v...)
		}
		if len(ans) >= k {
			return ans[:k]
		}
	}
	return ans
}

// @lc code=end
