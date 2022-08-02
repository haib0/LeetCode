package leetcode

import (
	"container/heap"
	"math"
)

type State struct {
	id   int
	dist float64
}

type StateHeap []State

func (h StateHeap) Len() int {
	return len(h)
}

func (h StateHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}

func (h StateHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *StateHeap) Push(x interface{}) {
	*h = append(*h, x.(State))
}

func (h *StateHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type neighbor struct {
	id     int
	weight float64
}

func dijkstra(start int, graph [][]neighbor) []float64 {
	n := len(graph)

	var visited = make([]bool, n)
	var minDist = make([]float64, n)

	for i := 0; i < n; i++ {
		minDist[i] = math.MaxFloat64
	}

	h := &StateHeap{State{id: start, dist: 0}}
	minDist[start] = 0

	for h.Len() > 0 {
		p := heap.Pop(h).(State).id
		if visited[p] {
			continue
		}
		visited[p] = true

		for _, neighbor := range graph[p] {
			// neighbor[0]: id, neighbor[1]: dist
			d := minDist[p] + neighbor.weight
			if d < minDist[neighbor.id] {
				minDist[neighbor.id] = d
				heap.Push(h, State{
					id: neighbor.id, dist: d,
				})
			}
		}
	}

	return minDist
}
