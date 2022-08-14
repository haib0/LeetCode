/*
 * @lc app=leetcode.cn id=1514 lang=golang
 *
 * [1514] 概率最大的路径
 */

package leetcode

import (
	"container/heap"
)

// @lc code=start

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	graph := make([][]neighbor1514, n)

	for i, v := range edges {
		from, to := v[0], v[1]
		weight := succProb[i]
		graph[from] = append(graph[from], neighbor1514{id: to, weight: weight})
	}

	return dijkstra1514(start, end, graph)
}

type State1514 struct {
	id   int
	dist float64
}

type StateHeap1514 []State1514

func (h StateHeap1514) Len() int {
	return len(h)
}

func (h StateHeap1514) Less(i, j int) bool {
	return h[i].dist > h[j].dist
}

func (h StateHeap1514) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *StateHeap1514) Push(x interface{}) {
	*h = append(*h, x.(State1514))
}

func (h *StateHeap1514) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type neighbor1514 struct {
	id     int
	weight float64
}

func dijkstra1514(start, end int, graph [][]neighbor1514) float64 {
	n := len(graph)

	var visited = make([]bool, n)
	var minDist = make([]float64, n)

	for i := 0; i < n; i++ {
		minDist[i] = -1
	}

	h := &StateHeap1514{State1514{id: start, dist: 0}}
	minDist[start] = 0

	for h.Len() > 0 {
		p := heap.Pop(h).(State).id
		if visited[p] {
			continue
		}
		visited[p] = true
		if p == end {
			return minDist[p]
		}

		for _, neighbor1514 := range graph[p] {
			// neighbor1514[0]: id, neighbor1514[1]: dist
			d := minDist[p] * neighbor1514.weight
			if minDist[neighbor1514.id] == -1 {
				d = neighbor1514.weight
			}
			if minDist[neighbor1514.id] == -1 || d > minDist[neighbor1514.id] {
				minDist[neighbor1514.id] = d
				heap.Push(h, State{
					id: neighbor1514.id, dist: d,
				})
			}
		}
	}

	return 0
}

// @lc code=end
