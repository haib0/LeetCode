/*
 * @lc app=leetcode.cn id=327 lang=golang
 *
 * [327] 区间和的个数
 */
package leetcode

// @lc code=start
var Count int
var temp []int
var Lower, Upper int

func countRangeSum(nums []int, lower int, upper int) int {
	Lower = lower
	Upper = upper
	Count = 0

	n := len(nums)
	preSum := make([]int, n+1)
	s := 0
	for i, v := range nums {
		s += v
		preSum[i+1] = s
	}

	temp = make([]int, n+1)
	sort327(preSum, 0, n)
	return Count
}

func sort327(nums []int, lo, hi int) {
	if lo == hi {
		return
	}
	mid := (lo + hi) / 2
	sort327(nums, lo, mid)
	sort327(nums, mid+1, hi)
	merge327(nums, lo, mid, hi)
}

func merge327(nums []int, lo, mid, hi int) {
	// copy(temp, nums)
	for p := lo; p <= hi; p++ {
		temp[p] = nums[p]
	}
	start, end := mid+1, mid+1
	for p := lo; p <= mid; p++ {
		for start <= hi && nums[start]-nums[p] < Lower {
			start++
		}
		for end <= hi && nums[end]-nums[p] <= Upper {
			end++
		}
		Count += end - start
	}

	i, j := lo, mid+1
	for p := lo; p <= hi; p++ {
		if i == mid+1 {
			nums[p] = temp[j]
			j++
		} else if j == hi+1 {
			nums[p] = temp[i]
			i++
		} else if temp[i] < temp[j] {
			nums[p] = temp[i]
			i++
		} else {
			nums[p] = temp[j]
			j++
		}
	}
}

// @lc code=end
