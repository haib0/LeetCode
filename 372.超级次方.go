/*
 * @lc app=leetcode.cn id=372 lang=golang
 *
 * [372] 超级次方
 */
package leetcode

// @lc code=start

const BASE = 1337

// -> a 的 k 次方 % BASE
func p372(a, k int) int {
	// (a * b) % k = (a % k)(b % k) % k
	a %= BASE
	ans := 1
	for i := 0; i < k; i++ {
		ans *= a
		ans %= BASE
	}
	return ans
}

// superPow(a, [1,2,0])
// = p(a, 0) * p(superPow([1,2]), 10) % BASE
func superPow(a int, b []int) int {
	l := len(b)
	if l == 1 {
		return p372(a, b[0])
	}
	x1 := p372(a, b[l-1])
	x2 := p372(superPow(a, b[:l-1]), 10)
	return x1 * x2 % BASE
}

// @lc code=end
