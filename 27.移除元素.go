/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */
package leetcode

// @lc code=start
func removeElement_(nums []int, val int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] != val {
			i++
			continue
		}

		nums[i], nums[j] = nums[j], nums[i]
		j--
	}
	return i
}

func removeElement(nums []int, val int) int {
	low, fast := 0, 0

	for fast < len(nums) {
		if nums[fast] != val {
			nums[low] = nums[fast]
			low++
		}
		fast++
	}
	return low
}

// @lc code=end
