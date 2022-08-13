/*
 * @lc app=leetcode.cn id=398 lang=golang
 *
 * [398] 随机数索引
 */
package leetcode

import "math/rand"

// @lc code=start
type Solution398 struct {
	nm map[int][]int
}

func Constructor398(nums []int) Solution398 {
	nm := make(map[int][]int)
	for i, v := range nums {
		nm[v] = append(nm[v], i)
	}
	return Solution398{nm: nm}
}

func (this *Solution398) Pick(target int) int {
	l := len(this.nm[target])
	return this.nm[target][rand.Intn(l)]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
// @lc code=end
