/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] O(1) 时间插入、删除和获取随机元素
 */
package leetcode

import "math/rand"

// @lc code=start
type RandomizedSet struct {
	list       []int
	valToIndex map[int]int
}

func Constructor380() RandomizedSet {
	return RandomizedSet{
		valToIndex: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.valToIndex[val]
	if ok {
		return false
	}
	this.valToIndex[val] = len(this.list)
	this.list = append(this.list, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.valToIndex[val]
	if !ok {
		return false
	}
	x := this.list[len(this.list)-1]
	this.valToIndex[x] = index
	this.list[index] = x
	this.list = this.list[:len(this.list)-1]
	delete(this.valToIndex, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	i := rand.Intn(len(this.list))
	return this.list[i]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
