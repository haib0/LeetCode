/*
 * @lc app=leetcode.cn id=563 lang=golang
 *
 * [563] 二叉树的坡度
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
func findTilt(root *TreeNode) (ans int) {
	var abs func(int) int
	abs = func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}

	var dfs func(*TreeNode) int
	dfs = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		leftSum := dfs(tn.Left)
		rightSum := dfs(tn.Right)
		ans += abs(leftSum - rightSum)
		return leftSum + rightSum + tn.Val
	}

	dfs(root)
	return
}

// @lc code=end
