/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
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

func reverseKGroup(head *ListNode, k int) *ListNode {
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}

	newHead := reverse(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead

}

// reverse [a, b)
func reverse(a, b *ListNode) *ListNode {
	pre, curr := &ListNode{}, a

	for curr != b {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}

	return pre
}

// @lc code=end
