package leetcode

import "sort"

/*
想象一下你是个城市基建规划者，地图上有 N 座城市，它们按以 1 到 N 的次序编号。
给你一些可连接的选项 conections，其中每个选项 conections[i] = [city1, city2, cost] 表示
将城市 city1 和城市 city2 连接所要的成本为 cost
（连接是双向的，也就是说城市 city1 和城市 city2 相连也同样意味着城市 city2 和城市 city1 相连）。
计算使得每对城市都连通的最小成本。如果根据已知条件无法完成该项任务，则请你返回 -1。

示例：
输入：N = 3, conections = [[1,2,5],[1,3,6],[2,3,1]]
输出：6
解释： 选出任意 2 条边都可以连接所有城市，我们从中选取成本最小的 2 条。
*/
func minimumCost(n int, conections [][]int) int {
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
