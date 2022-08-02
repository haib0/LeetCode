/*
 * @lc app=leetcode.cn id=698 lang=golang
 *
 * [698] 划分为k个相等的子集
 */
package leetcode

import "fmt"

// @lc code=start
func canPartitionKSubsets(nums []int, k int) bool {
	type b698 struct {
		nums []int
		sum  int
	}
	check := func(bucket []b698) bool {
		if len(bucket) == 0 {
			return true
		}

		s := bucket[0].sum
		for _, v := range bucket {
			if len(v.nums) > 0 && v.sum == s {
				continue
			}
			return false
		}

		return true
	}

	var ans bool
	n := len(nums)
	var bucket = make([]b698, k)
	var backtracking func(startNum int)
	
	backtracking = func(startNum int) {
		if ans {
			return
		}

		if startNum == n {
			x := check(bucket)
			if x {
				fmt.Println(bucket)
				ans = x
			}
			return
		}

		for j := 0; j < k; j++ {
			bucket[j].nums = append(bucket[j].nums, nums[startNum])
			bucket[j].sum += nums[startNum]

			backtracking(startNum + 1)

			bucket[j].nums = bucket[j].nums[:len(bucket[j].nums)-1]
			bucket[j].sum -= nums[startNum]
		}

	}

	backtracking(0)

	return ans
}

// @lc code=end
