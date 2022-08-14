/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
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
func rob(root *TreeNode) int {
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	var dfs func(node *TreeNode) (dorob, dontrob int)
	dfs = func(node *TreeNode) (dorob int, dontrob int) {
		if node == nil {
			return 0, 0
		}
		l, ln := dfs(node.Left)
		r, rn := dfs(node.Right)
		dorob = node.Val + ln + rn
		dontrob = max(l, ln) + max(r, rn)
		return
	}

	return max(dfs(root))
}

// @lc code=end
