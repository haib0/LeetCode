/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
 */

package leetcode

import (
	"math"
)

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nLayer := 0
	pre := []*TreeNode{root}
	for len(pre) > 0 {
		curr := []*TreeNode{}

		for _, node := range pre {
			if node.Left != nil {
				curr = append(curr, node.Left)
			}
			if node.Right != nil {
				curr = append(curr, node.Right)
			}
		}

		if len(curr) == 0 {
			return int(math.Pow(2, float64(nLayer))) - 1 + len(pre)
		}

		nLayer++
		pre = curr
	}

	return -1
}

// @lc code=end
