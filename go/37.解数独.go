/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

package leetcode

// @lc code=start

func solveSudoku(board [][]byte) {
	isValid := func(x, y int, num byte) bool {
		for i := 0; i < 9; i++ {
			if board[x][i] == num {
				return false
			}

			if board[i][y] == num {
				return false
			}

			if board[(x/3)*3+i/3][(y/3)*3+i%3] == num {
				return false
			}
		}
		return true
	}

	m, n := len(board), len(board[0])

	var backtracking func(x, y int) bool
	backtracking = func(x, y int) bool {
		if y == n {
			return backtracking(x+1, 0)
		}

		if x == m {
			return true
		}

		if board[x][y] != '.' {
			return backtracking(x, y+1)
		}

		for i := '1'; i <= '9'; i++ {
			a := byte(i)

			if !isValid(x, y, a) {
				continue
			}

			board[x][y] = a

			if backtracking(x, y+1) {
				return true
			}

			board[x][y] = '.'
		}

		return false
	}

	backtracking(0, 0)
}

// @lc code=end
