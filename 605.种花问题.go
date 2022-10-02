/*
 * @lc app=leetcode.cn id=605 lang=golang
 *
 * [605] 种花问题
 */
package leetcode

// @lc code=start

func canPlaceFlowers(flowerbed []int, n int) bool {
	var max int
	for i, v := range flowerbed {
		if v == 1 {
			continue
		}
		if i == 0 || flowerbed[i-1] == 0 {
			if i == len(flowerbed)-1 || flowerbed[i+1] == 0 {
				max++
				flowerbed[i] = 1
			}
		}
	}
	return n <= max
}

// @lc code=end
