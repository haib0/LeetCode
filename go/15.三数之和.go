/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
package leetcode

import "sort"

// @lc code=start
func threeSum(nums []int) [][]int {
	var ans [][]int
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	i := 0
	for i < len(nums) {
		v := nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			nl, nr := nums[left], nums[right]
			sum := v + nl + nr
			if sum < 0 {
				for left < right && nums[left] == nl {
					left++
				}
			} else if sum > 0 {
				for left < right && nums[right] == nr {
					right--
				}
			} else {
				ans = append(ans, []int{v, nl, nr})
				for left < right && nums[left] == nl {
					left++
				}
				for left < right && nums[right] == nr {
					right--
				}
			}
		}

		for i < len(nums) && nums[i] == v {
			i++
		}
	}

	return ans
}

// @lc code=end
