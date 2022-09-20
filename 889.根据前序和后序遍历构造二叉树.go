/*
 * @lc app=leetcode.cn id=889 lang=golang
 *
 * [889] 根据前序和后序遍历构造二叉树
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
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}

	v1 := preorder[1]
	l := 0
	for i, v2 := range postorder {
		if v2 == v1 {
			l = i + 1
			break
		}
	}

	root.Left = constructFromPrePost(preorder[1:l+1], postorder[:l])
	root.Right = constructFromPrePost(preorder[l+1:], postorder[l:len(postorder)-1])

	return root
}

// @lc code=end
