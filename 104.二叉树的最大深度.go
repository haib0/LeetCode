/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)

	if leftMax < rightMax {
		return 1 + rightMax
	}
	return 1 + leftMax
}

func maxDepthT(root *TreeNode) int {
	var ans, depth int

	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		depth++

		if depth > ans {
			ans = depth
		}

		traverse(node.Left)
		traverse(node.Right)

		depth--
	}

	traverse(root)
	return ans
}

func maxDepthBST(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 0
	pre := []*TreeNode{root}
	for len(pre) > 0 {
		result++
		curr := []*TreeNode{}
		for _, n := range pre {
			if n.Left != nil {
				curr = append(curr, n.Left)
			}
			if n.Right != nil {
				curr = append(curr, n.Right)
			}
		}
		pre = curr
	}
	return result
}

// @lc code=end
