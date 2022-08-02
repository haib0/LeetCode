/*
 * @lc app=leetcode.cn id=895 lang=golang
 *
 * [895] 最大频率栈
 */
package leetcode

// @lc code=start
type FreqStack struct {
	maxFreq     int
	valToFreq   map[int]int
	freqToStack map[int][]int
}

func Constructor895() FreqStack {
	return FreqStack{
		valToFreq:   make(map[int]int),
		freqToStack: make(map[int][]int),
	}
}

func (this *FreqStack) Push(val int) {
	freq := this.valToFreq[val]
	this.freqToStack[freq+1] = append(this.freqToStack[freq+1], val)
	this.valToFreq[val] = freq + 1
	if freq+1 > this.maxFreq {
		this.maxFreq = freq + 1
	}
}

func (this *FreqStack) Pop() int {
	n := len(this.freqToStack[this.maxFreq])
	x := this.freqToStack[this.maxFreq][n-1]
	this.freqToStack[this.maxFreq] = this.freqToStack[this.maxFreq][:n-1]
	this.valToFreq[x]--
	for this.maxFreq > 0 && len(this.freqToStack[this.maxFreq]) == 0 {
		this.maxFreq--
	}

	return x
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
// @lc code=end
