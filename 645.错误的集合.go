/*
 * @lc app=leetcode.cn id=645 lang=golang
 *
 * [645] 错误的集合
 */
package leetcode

// @lc code=start

func findErrorNums(nums []int) []int {
	var AbsInt = func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	var dup, missing int
	for _, v := range nums {
		i := AbsInt(v) - 1
		// 如果 num[i] < 0, 则表示 i+1 已出现过
		if nums[i] < 0 {
			dup = i + 1
		} else {
			nums[i] = -nums[i]
		}
	}
	for i, v := range nums {
		if v > 0 {
			missing = i + 1
		}
	}
	return []int{dup, missing}
}

// @lc code=end
