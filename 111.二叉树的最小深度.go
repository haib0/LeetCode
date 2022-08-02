/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ans = 1
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ { // BFS
			if queue[i].Left == nil && queue[i].Right == nil {
				return ans
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]
		ans++
	}
	return ans
}

func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lMin := minDepth(root.Left)
	rMin := minDepth(root.Right)

	if lMin == 0 {
		return rMin + 1
	}
	if rMin == 0 {
		return lMin + 1
	}

	if lMin < rMin {
		return lMin + 1
	}

	return rMin + 1
}

// @lc code=end
