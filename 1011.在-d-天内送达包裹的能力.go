/*
 * @lc app=leetcode.cn id=1011 lang=golang
 *
 * [1011] 在 D 天内送达包裹的能力
 */
package leetcode

// @lc code=start
func shipWithinDays(weights []int, days int) int {
	left, right := 0, int(1e8)

	for left < right {
		mid := (left + right) / 2
		if mid == 0 {
			return 1
		}
		if f1011(mid, weights) <= days {
			right = mid
		} else {
			left = mid + 1
		}
	}

	for _, c := range weights {
		if left < c {
			left = c
		}
	}
	return left
}

func f1011(carry int, weights []int) (days int) {
	tmp := 0
	for _, w := range weights {
		tmp += w
		if tmp > carry {
			days++
			tmp = w
		}
		if tmp == carry {
			days++
			tmp = 0
		}
	}
	if tmp > 0 {
		days++
	}
	return
}

// @lc code=end
