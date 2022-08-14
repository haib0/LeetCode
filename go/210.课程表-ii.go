/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

package leetcode

// @lc code=start

func findOrder(numCourses int, prerequisites [][]int) []int {
	g := buildGraph210(numCourses, prerequisites)

	h, p := hasCircle210(g)

	if h {
		return []int{}
	}

	for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}

	return p
}

func buildGraph210(n int, prerequisites [][]int) (graph [][]int) {
	graph = make([][]int, n)

	for _, nb := range prerequisites {
		a, b := nb[0], nb[1]
		graph[b] = append(graph[b], a)
	}

	return
}

func hasCircle210(graph [][]int) (has bool, postOrder []int) {
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
		postOrder = append(postOrder, x)
	}

	for i := 0; i < n; i++ {
		if !viewed[i] {
			traverse(i)
		}
	}

	return
}

// @lc code=end
