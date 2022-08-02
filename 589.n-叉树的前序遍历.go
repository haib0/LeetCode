/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N 叉树的前序遍历
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

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	ans := []int{root.Val}
	for _, child := range root.Children {
		r := preorder(child)
		ans = append(ans, r...)
	}
	return ans
}

// @lc code=end
