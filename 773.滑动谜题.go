/*
 * @lc app=leetcode.cn id=773 lang=golang
 *
 * [773] 滑动谜题
 */
package leetcode

import (
	"strconv"
)

// @lc code=start
func slidingPuzzle(board [][]int) int {
	neighbers := make(map[int][]int)
	neighbers[0] = []int{1, 3}
	neighbers[1] = []int{0, 2, 4}
	neighbers[2] = []int{1, 5}
	neighbers[3] = []int{0, 4}
	neighbers[4] = []int{1, 3, 5}
	neighbers[5] = []int{2, 4}

	boards := ""
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			boards += strconv.Itoa(board[i][j])
		}
	}

	var step int
	visited := make(map[string]bool)
	queue := []string{boards}
	visited[boards] = true
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			currBoard := queue[i]
			if currBoard == "123450" {
				return step
			}

			p0 := 0
			for ; currBoard[p0] != '0'; p0++ {
			}

			for _, pn := range neighbers[p0] {
				cu := []rune(currBoard)
				cu[p0], cu[pn] = cu[pn], cu[p0]
				nextBoard := string(cu)

				if !visited[nextBoard] {
					queue = append(queue, nextBoard)
					visited[nextBoard] = true
				}
			}
		}
		step++
		queue = queue[n:]
	}

	return -1
}

// @lc code=end
