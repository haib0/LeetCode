/*
 * @lc app=leetcode.cn id=1109 lang=golang
 *
 * [1109] 航班预订统计
 */
package leetcode

// @lc code=start
func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n)

	for _, booking := range bookings {
		first, last := booking[0]-1, booking[1]-1
		seats := booking[2]

		diff[first] += seats
		if last+1 < n {
			diff[last+1] -= seats
		}
	}

	ans := make([]int, n)
	tmp := 0
	for i, v := range diff {
		tmp += v
		ans[i] = tmp
	}

	return ans
}

// @lc code=end
