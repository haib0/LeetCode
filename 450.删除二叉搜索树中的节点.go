/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
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

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Right == nil {
			return root.Left
		}
		if root.Left == nil {
			return root.Right
		}

		min := findMinBSTNode(root.Right)
		root.Right = deleteNode(root.Right, min.Val)

		min.Left = root.Left
		min.Right = root.Right
		root = min
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}

	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}

func findMinBSTNode(root *TreeNode) *TreeNode {
	curr := root
	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}
	return curr
}

// @lc code=end
