/*
 * @lc app=leetcode.cn id=710 lang=golang
 *
 * [710] 黑名单中的随机数
 */
package leetcode

import "math/rand"

// @lc code=start
type Solution710 struct {
	size int
	mp   map[int]int
}

func Constructor710(n int, blacklist []int) Solution710 {
	size := n - len(blacklist)
	mp := make(map[int]int)
	for _, v := range blacklist {
		mp[v] = -1
	}

	last := n - 1
	for _, v := range blacklist {
		if v >= size {
			continue
		}
		for {
			if _, ok := mp[last]; ok {
				last--
				continue
			}
			break
		}
		mp[v] = last
		last--
	}
	return Solution710{
		size: size,
		mp:   mp,
	}
}

func (this *Solution710) Pick() int {
	x := rand.Intn(this.size)
	if v, ok := this.mp[x]; ok {
		return v
	}
	return x
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
// @lc code=end
