package leetcode

import (
	"testing"
)

/*
输入：N = 3, conections = [[1,2,5],[1,3,6],[2,3,1]]
输出：6
解释： 选出任意 2 条边都可以连接所有城市，我们从中选取成本最小的 2 条。
*/
func TestMinimumCost(t *testing.T) {
	n := 3
	connections := [][]int{{1, 2, 5}, {1, 3, 6}, {2, 3, 1}}

	expect := 6

	ans := minimumCost(n, connections)

	if ans != expect {
		t.Error(ans)
	}
}
