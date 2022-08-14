/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

package leetcode

// @lc code=start

func permute(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	var backtracking func(start int, track []int)
	backtracking = func(start int, track []int) {
		if len(track) == n {
			trackCopy := make([]int, n)
			copy(trackCopy, track)
			ans = append(ans, trackCopy)
			return
		}

		for i := 0; i < n; i++ {
			if SliceContain(track, nums[i]) {
				continue
			}

			track = append(track, nums[i])
			backtracking(i, track)
			track = track[:len(track)-1]
		}

	}

	backtracking(0, []int{})

	return ans
}

func SliceContain(s []int, x int) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}

	return false
}

// @lc code=end
