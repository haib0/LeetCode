/*
 * @lc app=leetcode.cn id=1094 lang=golang
 *
 * [1094] 拼车
 */
package leetcode

// @lc code=start
func carPooling(trips [][]int, capacity int) bool {
	diff := make([]int, 1000)

	for _, trip := range trips {
		n := trip[0]
		from, to := trip[1], trip[2]

		if n > capacity {
			return false
		}

		diff[from] += n
		if to+1 <= 1000 {
			diff[to] -= n
		}
	}

	need := 0
	for _, v := range diff {
		need += v
		if need > capacity {
			return false
		}
	}
	return true
}

// @lc code=end
