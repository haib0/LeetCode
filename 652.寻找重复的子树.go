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

var ans652 []*TreeNode
var nodeFreq map[string]int

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	ans652 = []*TreeNode{}
	nodeFreq = make(map[string]int)

	traverse652(root)

	return ans652
}

// 序列化
func traverse652(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	l := traverse652(root.Left)
	r := traverse652(root.Right)

	subTree := l + "," + r + "," + strconv.Itoa(root.Val)

	freq := nodeFreq[subTree]
	if freq == 1 {
		ans652 = append(ans652, root)
	}
	nodeFreq[subTree] = freq + 1

	return subTree
}

// @lc code=end
