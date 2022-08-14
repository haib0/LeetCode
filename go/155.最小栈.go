/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */
package leetcode

// @lc code=start
type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor155() MinStack {
	intSize := 32 << (^uint(0) >> 63) // 32 or 64
	MaxInt := 1<<(intSize-1) - 1
	return MinStack{
		[]int{},
		[]int{MaxInt},
	}
}

func (this *MinStack) Push(val int) {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	this.stack = append(this.stack, val)
	this.minStack = append(this.minStack, min(this.minStack[len(this.minStack)-1], val))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
