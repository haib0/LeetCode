/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	p := findFromEnd19(dummy, n+1)
	p.Next = p.Next.Next

	return dummy.Next
}

func findFromEnd19(head *ListNode, n int) *ListNode {
	p := head
	for i := 0; i < n; i++ {
		p = p.Next
	}
	pn := head
	for p != nil {
		p = p.Next
		pn = pn.Next
	}
	return pn
}

// @lc code=end
