/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	node := root.Left
	if node == nil {
		return
	}

	for node.Right != nil {
		node = node.Right
	}

	node.Right = root.Right
	root.Right = root.Left
	root.Left = nil
}

// @lc code=end
