/*
 * @lc app=leetcode.cn id=965 lang=golang
 *
 * [965] 单值二叉树
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

var vCache965 int

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	vCache965 = root.Val

	return isUnivalTreeT(root)
}

func isUnivalTreeT(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Val != vCache965 {
		return false
	}

	return isUnivalTreeT(root.Left) && isUnivalTreeT(root.Right)
}

// @lc code=end
