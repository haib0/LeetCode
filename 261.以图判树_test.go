package leetcode

import "testing"

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
func TestValidTree(t *testing.T) {
	n := 5

	edges1 := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}
	expect1 := true

	edges2 := [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}
	expect2 := false

	ans1 := validTree(n, edges1)
	ans2 := validTree(n, edges2)

	if ans1 != expect1 {
		t.Errorf("ans1 error: %v", ans1)
	}

	if ans2 != expect2 {
		t.Errorf("ans2 error: %v", ans2)
	}
}
