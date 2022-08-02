/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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

func buildTreeInPost(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: postorder[n-1],
	}

	idx := 0
	for inorder[idx] != root.Val {
		idx++
	}

	leftIn := inorder[:idx]
	leftPost := postorder[:idx]

	rightIn := inorder[idx+1:]
	rightPost := postorder[idx : n-1]

	root.Left = buildTree(leftIn, leftPost)
	root.Right = buildTree(rightIn, rightPost)

	return root
}

// @lc code=end
