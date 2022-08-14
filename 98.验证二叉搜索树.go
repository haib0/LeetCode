/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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

func isValidBST(root *TreeNode) bool {
	return isValidBST1(root, nil, nil)
}
func isValidBST1(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}

	if max != nil && root.Val >= max.Val {
		return false
	}

	if min != nil && root.Val <= min.Val {
		return false
	}

	return isValidBST1(root.Left, min, root) && isValidBST1(root.Right, root, max)
}

// @lc code=end
