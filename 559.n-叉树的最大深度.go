/*
 * @lc app=leetcode.cn id=559 lang=golang
 *
 * [559] N 叉树的最大深度
 */
package leetcode

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func maxDepthN(root *Node) int {
	if root == nil {
		return 0
	}

	depth := 0
	for _, child := range root.Children {
		d := maxDepthN(child)
		if depth < d {
			depth = d
		}
	}

	return depth + 1
}

// @lc code=end
