/*
 * @lc app=leetcode.cn id=517 lang=golang
 *
 * [517] 超级洗衣机
 */
package leetcode

// @lc code=start
func findMinMoves(machines []int) int {
	n := len(machines)

	tot := 0
	for _, i := range machines {
		tot += i
	}
	if tot%n != 0 {
		return -1
	}

	avg := tot / n
	ans := 0
	sum := 0
	for _, num := range machines {
		num_ := num - avg
		sum += num_
		sum_abs := sum
		if sum < 0 {
			sum_abs = -sum
		}
		if ans < sum_abs {
			ans = sum_abs
		}
		if ans < num_ {
			ans = num_
		}
	}
	return ans
}

// @lc code=end
