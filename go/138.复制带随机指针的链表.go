/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */
package leetcode

// @lc code=start
type Node138 struct {
	Val    int
	Next   *Node138
	Random *Node138
}

var cachedNode = map[*Node138]*Node138{}

func deepCopy(node *Node138) *Node138 {
	if node == nil {
		return nil
	}

	if n, has := cachedNode[node]; has {
		return n
	}

	newNode := &Node138{Val: node.Val}
	// 注意：要在深拷贝next和random之前
	cachedNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode

}

func copyRandomList(head *Node138) *Node138 {
	return deepCopy(head)
}

// @lc code=end
