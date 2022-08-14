/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 */
package leetcode

import (
	"sort"
)

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		e1, e2 := envelopes[i], envelopes[j]
		return e1[0] < e2[0] || (e1[0] == e2[0] && e1[1] > e2[1])
	})

	h := make([]int, len(envelopes))
	for i, e := range envelopes {
		h[i] = e[1]
	}
	return lenOfLIS(h)
}

func lenOfLIS(nums []int) int {
	l := 0
	piles := make([]int, len(nums))

	for _, n := range nums {
		// binary search n in piles, left bound
		left, right := 0, l
		for left < right {
			mid := (left + right) / 2
			if piles[mid] < n {
				left = mid + 1
			} else {
				right = mid
			}
		}

		if left == l {
			l++
		}

		piles[left] = n
	}
	return l
}

// @lc code=end
