package leetcode

import (
	"testing"
)

func TestDijkstra(t *testing.T) {
	graph := [][]neighbor{
		{}, {{0, 1}, {2, 2}}, {{3, 1}}, {},
	}
	ans := dijkstra(1, graph)
	t.Log(ans)
}
