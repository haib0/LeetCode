/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */
package leetcode

// @lc code=start

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var result *ListNode
	if l1.Val < l2.Val {
		result = &ListNode{l1.Val, nil}
		l1 = l1.Next
	} else {
		result = &ListNode{l2.Val, nil}
		l2 = l2.Next
	}
	curr := result
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = &ListNode{l1.Val, nil}
			curr = curr.Next
			l1 = l1.Next
		} else {
			curr.Next = &ListNode{l2.Val, nil}
			curr = curr.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}

	return result
}

// @lc code=end
