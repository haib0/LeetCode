/*
 * @lc app=leetcode.cn id=1024 lang=golang
 *
 * [1024] 视频拼接
 */
package leetcode

import "sort"

// @lc code=start
func videoStitching(clips [][]int, time int) int {
	sort.Slice(clips, func(i, j int) bool {
		return clips[i][0] < clips[j][0]
	})

	ans := 0
	i := 0
	currR := 0
	nextR := 0
	for i < len(clips) && clips[i][0] <= currR {
		for i < len(clips) && clips[i][0] <= currR {
			if nextR < clips[i][1] {
				nextR = clips[i][1]
			}
			i++
		}
		currR = nextR
		ans++
		if currR >= time {
			return ans
		}
	}

	return -1
}

// @lc code=end
