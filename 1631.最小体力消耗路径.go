/*
 * @lc app=leetcode.cn id=1631 lang=golang
 *
 * [1631] 最小体力消耗路径
 */

package leetcode

import (
	"container/heap"
	"log"
	"math"
)

// @lc code=start

func minimumEffortPath(heights [][]int) int {

	m, n := len(heights), len(heights[0])
	var costs = make([][]int, m)
	var visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		x := make([]int, n)
		for i := 0; i < n; i++ {
			x[i] = math.MaxInt
		}
		costs[i] = x
		visited[i] = make([]bool, n)
	}

	var h = &StateHeap1631{}
	heap.Push(h, State1631{cost: 0})
	costs[0][0] = 0

	var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for h.Len() > 0 {
		curr := heap.Pop(h).(State1631)
		visited[curr.x][curr.y] = true

		for _, d := range directions {
			x, y := curr.x+d[0], curr.y+d[1]

			if x < 0 || x >= m || y < 0 || y >= n {
				continue
			}
			if visited[x][y] {
				continue
			}

			c := absInt1631(heights[x][y] - heights[curr.x][curr.y])
			if c < curr.cost {
				c = curr.cost
			}

			if c < costs[x][y] {
				costs[x][y] = c
				heap.Push(h, State1631{x: x, y: y, cost: c})
			}

			log.Println(x, y, c)
		}
	}

	log.Println(costs)
	return costs[m-1][n-1]
}

func absInt1631(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type State1631 struct {
	x, y int
	cost int
}

type StateHeap1631 []State1631

func (h StateHeap1631) Len() int {
	return len(h)
}

func (h StateHeap1631) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}

func (h StateHeap1631) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *StateHeap1631) Push(x interface{}) {
	*h = append(*h, x.(State1631))
}

func (h *StateHeap1631) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// @lc code=end
