/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

package leetcode

// @lc code=start

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	currA := headA
	currB := headB
	for currA != currB {
		if currA != nil {
			currA = currA.Next
		} else {
			currA = headB
		}
		if currB != nil {
			currB = currB.Next
		} else {
			currB = headA
		}
	}
	return currA
}

// @lc code=end
