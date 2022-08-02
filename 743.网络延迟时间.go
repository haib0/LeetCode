/*
 * @lc app=leetcode.cn id=743 lang=golang
 *
 * [743] 网络延迟时间
 */

package leetcode

import (
	"container/heap"
	"math"
)

// @lc code=start

func networkDelayTime(times [][]int, n int, k int) int {

	graph := make([][][]int, n)
	for _, v := range times {
		from, to, weight := v[0]-1, v[1]-1, v[2]
		graph[from] = append(graph[from], []int{to, weight})
	}

	dist := dijkstra743(k-1, graph)

	var ans int
	for _, v := range dist {
		if v == math.MaxInt {
			return -1
		}

		if ans < v {
			ans = v
		}
	}

	return ans
}

type State743 struct {
	id   int
	dist int
}

type StateHeap743 []State743

func (h StateHeap743) Len() int {
	return len(h)
}

func (h StateHeap743) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}

func (h StateHeap743) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *StateHeap743) Push(x interface{}) {
	*h = append(*h, x.(State743))
}

func (h *StateHeap743) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func dijkstra743(start int, graph [][][]int) []int {
	n := len(graph)

	var visited = make([]bool, n)
	var minDist = make([]int, n)

	for i := 0; i < n; i++ {
		minDist[i] = math.MaxInt
	}

	h := &StateHeap743{State743{id: start, dist: 0}}
	minDist[start] = 0

	for h.Len() > 0 {
		p := heap.Pop(h).(State743).id
		if visited[p] {
			continue
		}
		visited[p] = true

		for _, neighbor := range graph[p] {
			// neighbor[0]: id, neighbor[1]: dist
			d := minDist[p] + neighbor[1]
			if d < minDist[neighbor[0]] {
				minDist[neighbor[0]] = d
				heap.Push(h, State743{
					id: neighbor[0], dist: d,
				})
			}
		}
	}

	return minDist
}

// @lc code=end
