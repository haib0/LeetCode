/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}

	idx := 0
	for inorder[idx] != root.Val {
		idx++
	}

	leftPre := preorder[1 : idx+1]
	leftIn := inorder[:idx]
	rightPre := preorder[idx+1:]
	rightIn := inorder[idx+1:]

	root.Left = buildTree(leftPre, leftIn)
	root.Right = buildTree(rightPre, rightIn)

	return root
}

// @lc code=end
