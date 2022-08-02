/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
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

func generateTrees(n int) []*TreeNode {
	var ns []int
	for i := 1; i <= n; i++ {
		ns = append(ns, i)
	}

	return generateSTrees(ns)
}

func generateSTrees(ns []int) []*TreeNode {
	ans := []*TreeNode{}
	for i, n := range ns {
		ls := generateSTrees(ns[:i])
		rs := generateSTrees(ns[i+1:])
		if len(ls) == 0 && len(rs) == 0 {
			ans = append(ans, &TreeNode{
				Val:   n,
				Left:  nil,
				Right: nil,
			})
		} else if len(ls) == 0 {
			for _, r := range rs {
				ans = append(ans, &TreeNode{
					Val:   n,
					Left:  nil,
					Right: r,
				})
			}
		} else if len(rs) == 0 {
			for _, l := range ls {
				ans = append(ans, &TreeNode{
					Val:   n,
					Left:  l,
					Right: nil,
				})
			}
		} else {
			for _, l := range ls {
				for _, r := range rs {
					ans = append(ans, &TreeNode{
						Val:   n,
						Left:  l,
						Right: r,
					})
				}
			}
		}
	}

	return ans
}

// @lc code=end
