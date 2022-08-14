/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
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
func convertBST(root *TreeNode) *TreeNode {
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
