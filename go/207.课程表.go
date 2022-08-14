/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

package leetcode

// @lc code=start

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph207(numCourses, prerequisites)
	return !hasCircle207(graph)
}

func buildGraph207(n int, prerequisites [][]int) (graph [][]int) {
	graph = make([][]int, n)

	for _, nb := range prerequisites {
		a, b := nb[0], nb[1]
		graph[b] = append(graph[b], a)
	}

	return
}

// if graph has circle, return true
func hasCircle207(graph [][]int) (has bool) {
	n := len(graph)

	viewed := make([]bool, n)
	onPath := make([]bool, n)

	var traverse func(x int)
	traverse = func(x int) {
		if onPath[x] {
			has = true
			return
		}

		if viewed[x] || onPath[x] {
			return
		}

		viewed[x] = true
		onPath[x] = true

		for _, v := range graph[x] {
			traverse(v)
		}

		onPath[x] = false
	}

	for i := 0; i < n; i++ {
		if !viewed[i] {
			traverse(i)
		}
	}
	return
}

// @lc code=end
