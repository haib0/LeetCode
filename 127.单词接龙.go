/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */
package leetcode

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	var isNeighber = func(x, y string) bool {
		diff := 0
		for i := range x {
			if x[i] != y[i] {
				diff++
			}
		}
		return diff == 1
	}
	var ans int
	var queue = []string{beginWord}
	var marked = make(map[string]bool)
	marked[beginWord] = true
	for len(queue) > 0 {
		ans++
		n := len(queue)
		for i := 0; i < n; i++ {
			curr := queue[i]
			if curr == endWord {
				return ans
			}
			for _, w := range wordList {
				if marked[w] {
					continue
				}
				if isNeighber(curr, w) {
					queue = append(queue, w)
					marked[w] = true
				}
			}
		}
		queue = queue[n:]
	}
	return 0
}

// @lc code=end
