/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

package leetcode

// @lc code=start
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, word, i, j) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i, j int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	if board[i][j] == byte('0') {
		return false
	}
	if board[i][j] == word[0] {
		cache := board[i][j]
		board[i][j] = byte('0')
		if len(word) == 1 {
			return true
		} else {
			r := dfs(board, word[1:], i+1, j) || dfs(board, word[1:], i, j+1) || dfs(board, word[1:], i-1, j) || dfs(board, word[1:], i, j-1)
			board[i][j] = cache
			return r
		}
	}
	return false
}

// @lc code=end
