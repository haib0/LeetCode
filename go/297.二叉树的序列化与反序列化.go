/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
 */

package leetcode

import (
	"strconv"
	"strings"
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

type Codec struct {
	SEP  rune
	NULL rune
}

func Constructor297() Codec {
	return Codec{SEP: ',', NULL: '#'}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var b strings.Builder
	this.serializeS(root, &b)

	return b.String()
}

func (this *Codec) serializeS(root *TreeNode, b *strings.Builder) {
	if root == nil {
		b.WriteRune(this.NULL)
		b.WriteRune(this.SEP)
		return
	}

	b.WriteString(strconv.Itoa(root.Val))
	b.WriteRune(this.SEP)

	this.serializeS(root.Left, b)
	this.serializeS(root.Right, b)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	datas := strings.Split(data, ",")

	var build func() *TreeNode
	build = func() *TreeNode {
		if len(datas) == 0 {
			return nil
		}

		v := (datas)[0]

		datas = datas[1:]

		if v == "#" {
			return nil
		}

		x, err := strconv.Atoi(v)
		if err != nil {
			return nil
		}
		root := &TreeNode{Val: x}

		root.Left = build()
		root.Right = build()

		return root
	}

	return build()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end
