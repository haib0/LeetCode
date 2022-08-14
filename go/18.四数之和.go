/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */
package leetcode

import "sort"

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	var ans [][]int
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums)-3; {
		v := nums[i]
		ansi := threeSum18(nums[i+1:], target-v)
		for _, av := range ansi {
			ans = append(ans, append(av, nums[i]))
		}

		for i < len(nums) && nums[i] == v {
			i++
		}
	}
	return ans
}

func threeSum18(nums []int, target int) [][]int {
	var ans [][]int
	// sort.Slice(nums, func(i, j int) bool {
	// 	return nums[i] < nums[j]
	// })
	i := 0
	for i < len(nums) {
		v := nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			nl, nr := nums[left], nums[right]
			sum := v + nl + nr
			if sum < target {
				for left < right && nums[left] == nl {
					left++
				}
			} else if sum > target {
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
