/*
 * @lc app=leetcode.cn id=744 lang=golang
 *
 * [744] 寻找比目标字母大的最小字母
 */
package leetcode

// @lc code=start
func nextGreatestLetter(letters []byte, target byte) byte {
	lo, hi := 0, len(letters)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if target >= letters[mid] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	if lo > len(letters)-1 {
		return letters[0]
	}
	return letters[lo]
}

// @lc code=end
