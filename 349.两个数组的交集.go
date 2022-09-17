/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */
package leetcode

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	in1 := make(map[int]bool)
	for _, v := range nums1 {
		in1[v] = true
	}
	var ans []int
	for _, v := range nums2 {
		if in1[v] {
			ans = append(ans, v)
			delete(in1, v) // can't dup
		}
	}
	return ans
}

// @lc code=end
