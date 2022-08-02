/*
 * @lc app=leetcode.cn id=765 lang=golang
 *
 * [765] 情侣牵手
 */

package leetcode

// @lc code=start

func minSwapsCouples(row []int) int {
	n := len(row) / 2

	uf := NewUnionFind(n)
	for i := 0; i < n; i++ {
		uf.Union(row[2*i]/2, row[2*i+1]/2)
	}

	return n - uf.Count()
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
