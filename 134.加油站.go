/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */
package leetcode

// @lc code=start

// 贪心算法
func canCompleteCircuit(gas []int, cost []int) int {
	var start int

	sum := 0
	for i, v := range gas {
		sum += v - cost[i]
	}
	if sum < 0 {
		return -1
	}

	tank := 0
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}

	return start
}

// @lc code=end
