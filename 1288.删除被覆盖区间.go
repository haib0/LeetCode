/*
 * @lc app=leetcode.cn id=1288 lang=golang
 *
 * [1288] 删除被覆盖区间
 */
package leetcode

import "sort"

// @lc code=start
func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	nDel := 0
	l, r := intervals[0][0], intervals[0][1]
	for _, v := range intervals[1:] {
		// 被覆盖
		if v[0] >= l && v[1] <= r {
			nDel++
		}
		// 相交, 则合并
		if v[0] <= r && v[1] >= r {
			r = v[1]
		}
		// 不相交
		if v[0] > r {
			l = v[0]
			r = v[1]
		}

	}
	return len(intervals) - nDel
}

// @lc code=end
