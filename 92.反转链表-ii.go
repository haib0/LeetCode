/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */
package leetcode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	pa, b := dummy, dummy
	for i := 0; i < right; i++ {
		b = b.Next
		if i < left-1 {
			pa = pa.Next
		}
	}
	a := pa.Next
	bp := b.Next

	pre := &ListNode{}
	curr := a
	for curr != bp {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}

	pa.Next = b
	a.Next = bp

	return dummy.Next
}

// @lc code=end
