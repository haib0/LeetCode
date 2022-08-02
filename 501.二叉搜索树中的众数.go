/*
 * @lc app=leetcode.cn id=501 lang=golang
 *
 * [501] 二叉搜索树中的众数
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

var currCount, maxCount int
var ans501 []int
var prev *TreeNode

func findMode(root *TreeNode) []int {
	prev = nil
	traverse501(root)

	return ans501
}

func traverse501(root *TreeNode) {
	if root == nil {
		return
	}

	traverse501(root.Left)

	if prev == nil {
		currCount = 1
		maxCount = 1
		ans501 = []int{root.Val}
	} else {
		if prev.Val == root.Val {
			currCount++
		} else {
			currCount = 1
		}

		if currCount == maxCount {
			ans501 = append(ans501, root.Val)
		}

		if currCount > maxCount {
			ans501 = []int{root.Val}
			maxCount = currCount
		}
	}

	prev = root

	traverse501(root.Right)
}

// @lc code=end
