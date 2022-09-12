/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */
package leetcode

// @lc code=start
func intToRoman(num int) string {
	var (
		r0 = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
		r1 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
		r2 = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
		r3 = []string{"", "M", "MM", "MMM"}
	)
	return r3[num/1000] +
		r2[num%1000/100] +
		r1[num%100/10] +
		r0[num%10]
}

// @lc code=end
