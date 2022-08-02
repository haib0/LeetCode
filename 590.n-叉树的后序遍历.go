/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N 叉树的后序遍历
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

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}

	ans := []int{}
	for _, child := range root.Children {
		r := postorder(child)
		ans = append(ans, r...)
	}
	ans = append(ans, root.Val)

	return ans
}

// @lc code=end
