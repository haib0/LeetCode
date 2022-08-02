/*
 * @lc app=leetcode.cn id=1038 lang=golang
 *
 * [1038] 从二叉搜索树到更大和树
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

func bstToGst(root *TreeNode) *TreeNode {
	curr := 0
	var postOrder func(this *TreeNode)
	postOrder = func(this *TreeNode) {
		if this == nil {
			return
		}
		postOrder(this.Right)
		curr += this.Val
		this.Val = curr
		postOrder(this.Left)
	}

	postOrder(root)
	return root
}

// @lc code=end
