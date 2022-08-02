/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */
package leetcode

// @lc code=start
type MyQueue struct {
	I, O []int
}

func Constructor232() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.I = append(this.I, x)
}

func (this *MyQueue) Pop() int {
	peek := this.Peek()
	this.O = this.O[:len(this.O)-1]
	return peek
}

func (this *MyQueue) Peek() int {
	var peek int
	if len(this.O) == 0 {
		last := len(this.I) - 1
		for last >= 0 {
			this.O = append(this.O, this.I[last])
			last--
		}
		this.I = []int{}
	}
	peek = this.O[len(this.O)-1]
	return peek
}

func (this *MyQueue) Empty() bool {
	return len(this.I)+len(this.O) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
