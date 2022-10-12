/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 */

package leetcode

import "fmt"

// @lc code=start

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func binaryTreePaths(root *TreeNode) []string {
	var ans []string
	var backtrack func(root *TreeNode, path string)
	backtrack = func(root *TreeNode, path string) {
		if root == nil {
			return
		}
		path += fmt.Sprint(root.Val) + "->"
		if root.Left == nil && root.Right == nil {
			path = path[:len(path)-2]
			ans = append(ans, path)
			return
		}
		if root.Left != nil {
			backtrack(root.Left, path)
		}
		if root.Right != nil {
			backtrack(root.Right, path)
		}
	}

	backtrack(root, "")
	return ans
}

// @lc code=end
