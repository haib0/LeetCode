/*
 * @lc app=leetcode.cn id=652 lang=golang
 *
 * [652] 寻找重复的子树
 */

package leetcode

import (
	"strconv"
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

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var ans []*TreeNode
	freq := make(map[string]int)
	var traverse func(root *TreeNode) string
	traverse = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}
		l, r := traverse(root.Left), traverse(root.Right)
		x := l + "," + r + "," + strconv.Itoa(root.Val)
		if i := freq[x]; i == 1 {
			ans = append(ans, root)
		}
		freq[x]++
		return x
	}
	traverse(root)
	return ans
}

// @lc code=end
