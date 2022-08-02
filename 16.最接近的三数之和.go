/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */
package leetcode

import "sort"

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	ans := nums[0] + nums[1] + nums[2]
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i, v := range nums {
		left, right := i+1, len(nums)-1
		for left < right {
			nl, nr := nums[left], nums[right]
			sum := v + nl + nr
			var da int
			if sum < target {
				da = target - sum

				for left < right && nums[left] == nl {
					left++
				}
			} else if sum > target {
				da = sum - target
				for left < right && nums[right] == nr {
					right--
				}
			} else {
				return sum
			}

			if da < abs16(ans-target) {
				ans = sum
			}
		}
	}
	return ans
}

func abs16(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// @lc code=end
