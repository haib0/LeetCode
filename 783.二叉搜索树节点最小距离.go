/*
 * @lc app=leetcode.cn id=783 lang=golang
 *
 * [783] 二叉搜索树节点最小距离
 */

package leetcode

import "math"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	var minDiff = math.MaxInt
	var prev *TreeNode

	var traverse func(this *TreeNode)
	traverse = func(this *TreeNode) {
		if this == nil {
			return
		}

		traverse(this.Left)
		if prev != nil {
			diff := this.Val - prev.Val
			if diff < 0 {
				diff = -diff
			}

			if diff < minDiff {
				minDiff = diff
			}
		}
		prev = this
		traverse(this.Right)
	}

	traverse(root)
	return minDiff
}

// @lc code=end
