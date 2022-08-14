/*
 * @lc app=leetcode.cn id=797 lang=golang
 *
 * [797] 所有可能的路径
 */

package leetcode

// @lc code=start

func allPathsSourceTarget(graph [][]int) [][]int {
	var (
		ans      [][]int
		path     []int
		traverse func(x int)
	)

	n := len(graph)

	traverse = func(x int) {
		path = append(path, x)

		defer func() {
			path = path[0 : len(path)-1]
		}()

		if x == n-1 {
			// path slice need deep copy
			// copy need make dst
			p := make([]int, len(path))
			copy(p, path)
			ans = append(ans, p)
			return
		}

		for _, v := range graph[x] {

			traverse(v)
		}
	}

	traverse(0)
	return ans

}

// @lc code=end
