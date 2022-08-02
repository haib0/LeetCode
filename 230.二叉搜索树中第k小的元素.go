/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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

func kthSmallest(root *TreeNode, k int) int {
	ans := root.Val

	var preOrder func(*TreeNode)
	preOrder = func(this *TreeNode) {
		if this == nil {
			return
		}

		preOrder(this.Left)
		k--
		if k == 0 {
			ans = this.Val
		}
		preOrder(this.Right)
	}
	preOrder(root)
	return ans
}

// @lc code=end
