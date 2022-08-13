/*
 * @lc app=leetcode.cn id=382 lang=golang
 *
 * [382] 链表随机节点
 */
package leetcode

import "math/rand"

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution382 struct {
	Head *ListNode
}

func Constructor(head *ListNode) Solution382 {
	return Solution382{
		Head: head,
	}
}

func (this *Solution382) GetRandom() int {
	var ans int
	// 第 i 个数字, 选择的概率应为 1/i, 保持原选择的概率应为 1-1/i
	i := 0
	h := this.Head
	for h != nil {
		i++
		// 生成 [0, i] 中的随机数, 等于 0 的概率等于 1/i
		if rand.Intn(i) == 0 {
			ans = h.Val
		}
		h = h.Next
	}
	return ans
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
// @lc code=end
