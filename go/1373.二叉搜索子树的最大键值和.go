/*
 * @lc app=leetcode.cn id=1373 lang=golang
 *
 * [1373] 二叉搜索子树的最大键值和
 */

package leetcode

import (
	"math"
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

func maxSumBST(root *TreeNode) int {
	var ans int

	checkAns := func(sum int) {
		if ans < sum {
			ans = sum
		}
	}

	var traverse func(node *TreeNode) (isBST bool, minNode, maxNode int, sum int)
	traverse = func(node *TreeNode) (isBST bool, minNode int, maxNode int, sum int) {
		if node == nil {
			isBST = true
			minNode = math.MaxInt
			maxNode = math.MinInt
			sum = 0
			return
		}

		lIsBST, lMinN, lMaxN, lSum := traverse(node.Left)
		rIsBST, rMinN, rMaxN, rSum := traverse(node.Right)

		if lIsBST && rIsBST && node.Val > lMaxN && node.Val < rMinN {
			isBST = true

			if lMinN < node.Val {
				minNode = lMinN
			} else {
				minNode = node.Val
			}

			if rMaxN < node.Val {
				maxNode = node.Val
			} else {
				maxNode = rMaxN
			}

			sum = node.Val + lSum + rSum

			checkAns(sum)

			return

		}

		isBST = false
		return
	}

	traverse(root)
	return ans
}

func maxSumBST1(root *TreeNode) int {
	var isBST func(node *TreeNode) (ok bool, maxSum int)
	isBST = func(node *TreeNode) (ok bool, maxSum int) {
		if node == nil {
			return true, 0
		}

		lOk, lSum := isBST(node.Left)
		rOk, rSum := isBST(node.Right)

		if lOk && rOk {
			if node.Left != nil && node.Left.Val >= node.Val {
				if lSum < rSum {
					return false, rSum
				}
				return false, lSum
			}
			if node.Right != nil && node.Right.Val <= node.Val {
				if lSum < rSum {
					return false, rSum
				}
				return false, lSum
			}

			if node.Val < 0 {
				return true, lSum + rSum
			}
			return true, lSum + node.Val + rSum
		} else {
			if lSum < rSum {
				return false, rSum
			}
			return false, lSum
		}
	}

	_, ans := isBST(root)
	if ans < 0 {
		return 0
	}
	return ans
}

// @lc code=end
