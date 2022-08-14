/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */
package leetcode

import "container/heap"

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int {
	return len(h)
}

func (h ListNodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h ListNodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	ans := &ListNode{}
	p := ans

	h := &ListNodeHeap{}

	for _, x := range lists {
		if x != nil {
			heap.Push(h, x)
		}
	}

	for h.Len() != 0 {
		x := heap.Pop(h).(*ListNode)
		p.Next = x
		// no nil in heap
		if x.Next != nil {
			heap.Push(h, x.Next)
		}
		p = p.Next
	}

	return ans.Next
}

// @lc code=end
