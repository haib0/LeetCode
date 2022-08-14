/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] 用最少数量的箭引爆气球
 */
package leetcode

import "sort"

// @lc code=start
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	n := 0
	r := points[0][1]
	for _, p := range points[1:] {
		if p[0] <= r {
			n++
		} else {
			r = p[1]
		}
	}
	return len(points) - n
}

// @lc code=end
