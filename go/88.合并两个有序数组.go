/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */
package leetcode

// @lc code=start
func merge88(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1

	for p := m + n - 1; p >= 0; p-- {
		if p1 >= 0 && p2 >= 0 {
			if nums1[p1] >= nums2[p2] {
				nums1[p] = nums1[p1]
				p1--
			} else {
				nums1[p] = nums2[p2]
				p2--
			}
		} else if p2 >= 0 {
			nums1[p] = nums2[p2]
			p2--
		}

	}
}

// @lc code=end
