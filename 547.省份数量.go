/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 省份数量
 */
package leetcode

// @lc code=start
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	var visited = make([]bool, n)
	var dfs func(i int)
	dfs = func(i int) {
		visited[i] = true
		for j := 0; j < n; j++ {
			if !visited[j] && isConnected[i][j] == 1 {
				dfs(j)
			}
		}
	}

	var ans int
	for i := 0; i < n; i++ {
		if !visited[i] {
			ans++
			dfs(i)
		}
	}
	return ans
}

// @lc code=end
