/*
 * @lc app=leetcode.cn id=886 lang=golang
 *
 * [886] 可能的二分法
 */

package leetcode

// @lc code=start

func possibleBipartition(n int, dislikes [][]int) bool {
	graph := buildGraph886(n, dislikes)
	return isBipartite886(graph)
}

func buildGraph886(n int, neighbers [][]int) (graph [][]int) {
	graph = make([][]int, n)

	for _, nb := range neighbers {
		a, b := nb[0]-1, nb[1]-1
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	return
}

func isBipartite886(graph [][]int) bool {
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
