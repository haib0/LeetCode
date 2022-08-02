/*
 * @lc app=leetcode.cn id=1584 lang=golang
 *
 * [1584] 连接所有点的最小费用
 */

package leetcode

import "sort"

// @lc code=start
func minCostConnectPoints(points [][]int) int {
	var edges [][]int

	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			w := absInt1584(points[i][0]-points[j][0]) + absInt1584(points[i][1]-points[j][1])
			edges = append(edges, []int{i, j, w})
		}
	}

	return minimumCost1584(n, edges)
}

func absInt1584(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minimumCost1584(n int, conections [][]int) int {
	// Kruskal 最小生成树算法
	var mst int

	uf := NewUnionFind(n + 1)
	sort.Slice(conections, func(i, j int) bool {
		return conections[i][2] < conections[j][2]
	})

	for _, edge := range conections {
		u, v := edge[0], edge[1]

		if uf.IsConnected(u, v) {
			continue
		}

		uf.Union(u, v)
		mst += edge[2]
	}

	if uf.Count() != 2 {
		return -1
	}
	return mst
}

// type UnionFind struct {
// 	n      int   // 节点个数
// 	count  int   // 连通分量个数
// 	parent []int //父节点
// }

// func NewUnionFind(n int) *UnionFind {
// 	p := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		p[i] = i
// 	}

// 	return &UnionFind{
// 		n:      n,
// 		count:  n,
// 		parent: p,
// 	}
// }

// func (uf *UnionFind) Union(p, q int) {
// 	rp := uf.findRoot(p)
// 	rq := uf.findRoot(q)

// 	if rp == rq {
// 		return
// 	}

// 	uf.parent[rq] = rp
// 	uf.count--
// }

// func (uf *UnionFind) IsConnected(p, q int) bool {
// 	rp := uf.findRoot(p)
// 	rq := uf.findRoot(q)

// 	return rp == rq
// }

// func (uf *UnionFind) findRoot(x int) int {
// 	// 路径压缩
// 	if uf.parent[x] != x {
// 		uf.parent[x] = uf.findRoot(uf.parent[x])
// 	}
// 	return uf.parent[x]
// }

// func (uf *UnionFind) Count() int {
// 	return uf.count
// }

// @lc code=end
