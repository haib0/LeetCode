/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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
func maxPathSum(root *TreeNode) int {
	const MINVAL = -1000
	var maxInt = func(nums ...int) int {
		if len(nums) == 0 {
			panic(nil)
		}
		x := nums[0]
		for _, v := range nums[1:] {
			if v > x {
				x = v
			}
		}
		return x
	}

	var ans = MINVAL
	check := func(tmp int) {
		if ans < tmp {
			ans = tmp
		}
	}
	// the best branch's sum with root
	var dp func(root *TreeNode) int
	dp = func(root *TreeNode) int {
		if root == nil {
			return MINVAL
		}
		l, r := dp(root.Left), dp(root.Right)
		tmp := maxInt(root.Val, l+root.Val, r+root.Val)
		check(maxInt(tmp, l+r+root.Val))
		return tmp
	}
	dp(root)
	return ans
}

// @lc code=end
