/*
 * @lc app=leetcode.cn id=701 lang=golang
 *
 * [701] 二叉搜索树中的插入操作
 */

package leetcode

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
			return root
		} else {
			root.Left = insertIntoBST(root.Left, val)
			return root
		}
	}
	if val > root.Val {
		if root.Right == nil && val > root.Val {
			root.Right = &TreeNode{Val: val}
			return root
		} else {
			root.Right = insertIntoBST(root.Right, val)
			return root
		}
	}
	return root
}

// @lc code=end
