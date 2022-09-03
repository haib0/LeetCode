/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */
package leetcode

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)

	nums := make([]int, m+n)
	i, i1, i2 := 0, 0, 0
	for i < m+n {
		var c int
		if i1 == m {
			c = nums2[i2]
			i2++
		} else if i2 == n {
			c = nums1[i1]
			i1++
		} else {
			if nums1[i1] < nums2[i2] {
				c = nums1[i1]
				i1++
			} else {
				c = nums2[i2]
				i2++
			}
		}
		nums[i] = c
		i++
	}

	i = (m + n) / 2
	if (m+n)%2 == 0 {
		return float64(nums[i-1]+nums[i]) / 2
	}

	return float64(nums[i])
}

// @lc code=end
