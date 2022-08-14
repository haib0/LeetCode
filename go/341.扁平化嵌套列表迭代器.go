/*
 * @lc app=leetcode.cn id=341 lang=golang
 *
 * [341] 扁平化嵌套列表迭代器
 */
package leetcode

// @lc code=start

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	isVal bool
	val   int
	list  []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	return this.isVal
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	return this.val
}

// Set this NestedInteger to hold a single integer.
func (this *NestedInteger) SetInteger(value int) {
	this.val = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {
	if this.isVal {
		this.isVal = false
		vN := NestedInteger{val: this.val}
		this.list = []*NestedInteger{&vN}
	}

	this.list = append(this.list, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	return this.list
}

type NestedIterator struct {
	list []*NestedInteger
}

func Constructor341(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		list: nestedList,
	}
}

func (this *NestedIterator) Next() int {
	x := this.list[0].GetInteger()
	this.list = this.list[1:]
	return x
}

func (this *NestedIterator) HasNext() bool {
	for len(this.list) != 0 && !this.list[0].IsInteger() {
		newL := this.list[0].GetList()
		newL = append(newL, this.list[1:]...)
		this.list = newL
	}
	return len(this.list) != 0
}

// @lc code=end
