/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */

package leetcode

import "sort"

// @lc code=start

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	var ans [][]int
	for _, p := range people {
		k := p[1]
		ans = append(ans[:k], append([][]int{p}, ans[k:]...)...)
	}
	return ans
}

// @lc code=end
