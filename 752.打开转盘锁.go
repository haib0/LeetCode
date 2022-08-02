/*
 * @lc app=leetcode.cn id=752 lang=golang
 *
 * [752] 打开转盘锁
 */
package leetcode

// @lc code=start
func openLock(deadends []string, target string) int {
	deadendsM := make(map[string]bool)
	for _, v := range deadends {
		deadendsM[v] = true
	}

	plusOne := func(curr string, i int) string {
		x := []byte(curr)
		if x[i] == '9' {
			x[i] = '0'
		} else {
			x[i] = x[i] + 1
		}
		curr = string(x)
		return curr
	}
	minusOne := func(curr string, i int) string {
		x := []byte(curr)
		if x[i] == '0' {
			x[i] = '9'
		} else {
			x[i] = x[i] - 1
		}
		curr = string(x)
		return curr
	}

	// BFS
	var step int
	var queue []string
	visited := make(map[string]bool)

	queue = append(queue, "0000")
	for len(queue) > 0 {
		n := len(queue)
		for ii := 0; ii < n; ii++ {
			curr := queue[ii]
			visited[curr] = true
			if deadendsM[curr] {
				continue
			}

			if curr == target {
				return step
			}

			for i := 0; i < 4; i++ {
				nextPlus := plusOne(curr, i)
				if !visited[nextPlus] {
					queue = append(queue, nextPlus)
				}

				nextMinus := minusOne(curr, i)
				if !visited[nextMinus] {
					queue = append(queue, nextMinus)
				}
			}
		}
		step++
		queue = queue[n:]
	}
	return -1
}

// @lc code=end
