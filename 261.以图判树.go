package leetcode

/*
判断输入的若干条边是否能构造出一棵树结构
------
n = 5
edges = [[0,1], [0,2], [0,3], [1,4]]
return true
------
n = 5
edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]
return false
*/
func validTree(n int, edges [][]int) bool {
	uf := NewUnionFind(n)

	for _, v := range edges {
		x, y := v[0], v[1]

		if uf.IsConnected(x, y) {
			return false
		}

		uf.Union(x, y)
	}

	return uf.Count() == 1
}
