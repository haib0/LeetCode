/*
 * @lc app=leetcode.cn id=875 lang=golang
 *
 * [875] 爱吃香蕉的珂珂
 */
package leetcode

// @lc code=start
func minEatingSpeed(piles []int, h int) int {
	left, right := 0, piles[0]
	for _, i := range piles {
		if i > right {
			right = i
		}
	}

	for left < right {
		mid := (left + right) / 2
		if mid == 0 {
			return 1
		}
		if needHours(mid, piles) <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func needHours(speed int, piles []int) (hours int) {
	for _, i := range piles {
		hours += i / speed
		if i%speed > 0 {
			hours++
		}
	}
	return
}

// @lc code=end
