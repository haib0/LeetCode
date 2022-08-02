/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */
package leetcode

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	var q *MonoQueue = &MonoQueue{}
	var ans []int
	for i, v := range nums {
		if i < k-1 {
			q.push(v)
			continue
		}

		q.push(v)
		ans = append(ans, q.max())
		q.pop(nums[i-k+1])
	}

	return ans
}

type MonoQueue struct {
	q []int
}

func (queue *MonoQueue) push(n int) {
	for len(queue.q) > 0 && queue.q[len(queue.q)-1] < n {
		queue.q = queue.q[:len(queue.q)-1]
	}
	queue.q = append(queue.q, n)
}

func (queue *MonoQueue) pop(n int) {
	if queue.q[0] == n {
		queue.q = queue.q[1:]
	}
}

func (queue *MonoQueue) max() int {
	return queue.q[0]
}

// @lc code=end
