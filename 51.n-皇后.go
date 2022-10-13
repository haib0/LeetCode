/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

package leetcode

// @lc code=start

func solveNQueens(n int) [][]string {
	var ans [][]string

	trace := make([][]rune, n)
	for i := 0; i < n; i++ {
		trace[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			trace[i][j] = '.'
		}
	}

	rToS := func(xr [][]rune) []string {
		var xs []string
		for _, v := range xr {
			xs = append(xs, string(v))
		}
		return xs
	}

	var backtracking func(x int)
	backtracking = func(x int) {
		if x == n {
			ans = append(ans, rToS(trace))
			return
		}

		for i := 0; i < n; i++ {
			if !isQueenOk(trace, x, i) {
				continue
			}

			trace[x][i] = 'Q'

			backtracking(x + 1)

			trace[x][i] = '.'
		}
	}

	backtracking(0)

	return ans
}

func isQueenOk(chess [][]rune, x, y int) bool {
	// chess like [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
	n := len(chess)

	for i := 0; i < n; i++ {
		if chess[i][y] != '.' {
			return false
		}
		if chess[x][i] != '.' {
			return false
		}
		ltx, lty := x-i, y-i
		if ltx >= 0 && lty >= 0 && chess[ltx][lty] != '.' {
			return false
		}
		rtx, rty := x-i, y+i
		if rtx >= 0 && rty < n && chess[rtx][rty] != '.' {
			return false
		}
	}
	return true
}

// @lc code=end
