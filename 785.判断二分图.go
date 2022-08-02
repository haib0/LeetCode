/*
 * @lc app=leetcode.cn id=785 lang=golang
 *
 * [785] 判断二分图
 */

package leetcode

// @lc code=start

func isBipartite(graph [][]int) bool {
	var ans = true
	n := len(graph)
	var colored = make([]bool, n)
	var visited = make([]bool, n)

	var traverse func(x int)
	traverse = func(x int) {
		if !ans {
			return
		}

		visited[x] = true

		for _, v := range graph[x] {
			if visited[v] {
				if colored[v] == colored[x] {
					ans = false
				}
			} else {
				colored[v] = !colored[x]
				traverse(v)
			}
		}
	}

	// Multi children graph
	for i := 0; i < n; i++ {
		if !visited[i] {
			traverse(i)
		}
	}

	return ans
}

// @lc code=end
