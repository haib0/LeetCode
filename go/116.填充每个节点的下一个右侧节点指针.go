/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

package leetcode

type Node116 struct {
	Val   int
	Left  *Node116
	Right *Node116
	Next  *Node116
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node116) *Node116 {
	if root == nil {
		return nil
	}

	pre := []*Node116{root}
	for len(pre) > 0 {
		curr := []*Node116{}
		for i, n := range pre {
			if n.Left != nil {
				curr = append(curr, n.Left)
			}
			if n.Right != nil {
				curr = append(curr, n.Right)
			}

			if i == len(pre)-1 {
				n.Next = nil
			} else {
				n.Next = pre[i+1]
			}
		}

		pre = curr
	}
	return root
}

// @lc code=end
