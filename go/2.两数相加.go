/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy *ListNode = &ListNode{}

	carry := 0
	p1, p2 := l1, l2
	p := dummy
	for p1 != nil || p2 != nil || carry > 0 {
		var v1, v2 int
		if p1 != nil {
			v1 = p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			v2 = p2.Val
			p2 = p2.Next
		}

		s := v1 + v2 + carry
		if s > 9 {
			s = s - 10
			carry = 1
		} else {
			carry = 0
		}

		p.Next = &ListNode{
			Val: s,
		}
		p = p.Next
	}
	return dummy.Next
}

// @lc code=end
