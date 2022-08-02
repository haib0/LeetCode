/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */
package leetcode

// @lc code=start
type MyStack struct {
	queue []int
}

func Constructor225() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	n := len(this.queue)
	newQueue := make([]int, n-1)
	var top int
	for i, v := range this.queue {
		if i == n-1 {
			top = v
			break
		}
		newQueue[i] = v
	}
	this.queue = newQueue
	return top
}

func (this *MyStack) Top() int {
	n := len(this.queue)
	newQueue := make([]int, n)
	var top int
	for i, v := range this.queue {
		if i == n-1 {
			top = v
		}
		newQueue[i] = v
	}
	this.queue = newQueue
	return top
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end
