/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */
package leetcode

// @lc code=start

// 位运算: n&(n-1)
func hammingWeight(num uint32) int {
	var ans int
	for num != 0 {
		num &= num - 1
		ans++
	}
	return ans
}

// @lc code=end
