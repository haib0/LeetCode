/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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

var maxDiameter int

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter = 0
	maxDepth543(root)
	return maxDiameter
}

func maxDepth543(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMax := maxDepth543(root.Left)
	rightMax := maxDepth543(root.Right)

	if maxDiameter < leftMax+rightMax {
		maxDiameter = leftMax + rightMax
	}

	if leftMax < rightMax {
		return 1 + rightMax
	}
	return 1 + leftMax
}

// @lc code=end
