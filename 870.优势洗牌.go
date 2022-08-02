/*
 * @lc app=leetcode.cn id=870 lang=golang
 *
 * [870] 优势洗牌
 */
package leetcode

import "sort"

// @lc code=start
func advantageCount(nums1 []int, nums2 []int) []int {
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] > nums1[j]
	})

	var nums2p [][]int
	for i, v := range nums2 {
		nums2p = append(nums2p, []int{i, v})
	}
	sort.Slice(nums2p, func(i, j int) bool {
		return nums2p[i][1] > nums2p[j][1]
	})

	ans := make([]int, len(nums1))
	head, tail := 0, len(nums1) - 1

	for _, v := range nums2p {
		if nums1[head] > v[1] {
			ans[v[0]] = nums1[head]
			head++
		}else {
			ans[v[0]] = nums1[tail]
			tail--
		}
	}

	return ans
}

// @lc code=end
