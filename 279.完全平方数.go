/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */
package leetcode

// @lc code=start
func numSquares(n int) int {
	var geneSquares = func(n int) (squares []int) {
		var x = 1
		var diff = 3
		for x <= n {
			squares = append(squares, x)
			x += diff
			diff += 2
		}
		return
	}
	var squares = geneSquares(n)
	var ans int
	var queue = []int{n}
	var viewed = make([]bool, n+1)
	viewed[n] = true
	for len(queue) > 0 {
		ans++
		n := len(queue)
		for i := 0; i < n; i++ {
			curr := queue[i]
			for _, v := range squares {
				next := curr - v
				if next == 0 {
					return ans
				}
				if next < 0 || viewed[next] {
					continue
				}
				queue = append(queue, next)
				viewed[next] = true
			}
		}
		queue = queue[n:]
	}
	return n
}

// @lc code=end
