/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] 第一个错误的版本
 */
package leetcode

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	var isBadVersion func(version int) bool
	lo, hi := 1, n
	for lo < hi {
		mid := (lo + hi) / 2
		if isBadVersion(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// @lc code=end
