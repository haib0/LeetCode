/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树
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

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxV := nums[0]
	maxI := 0
	for i, v := range nums {
		if maxV < v {
			maxV = v
			maxI = i
		}
	}
	root := &TreeNode{Val: maxV}

	root.Left = constructMaximumBinaryTree(nums[:maxI])
	root.Right = constructMaximumBinaryTree(nums[maxI+1:])

	return root
}

// @lc code=end
