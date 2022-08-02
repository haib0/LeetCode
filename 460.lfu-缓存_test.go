package leetcode

import "testing"

func TestLFU(t *testing.T) {
	lfu := Constructor460(2)

	lfu.Put(1, 1)
	lfu.Put(2, 2)
	lfu.Get(1)
	lfu.Put(3, 3)
	lfu.Get(2)
	lfu.Get(3)
	lfu.Put(4, 4)
	lfu.Get(1)
	lfu.Get(3)
	lfu.Get(4)
}

var node1 = &DNode{
	key: 1,
	val: 1,
}

var node2 = &DNode{
	key: 2,
	val: 2,
}

func TestDList(t *testing.T) {
	dl := NewDoublyLinkedList()

	dl.AddLast(node1)
	dl.AddLast(node2)
	dl.Remove(node1)
}

func TestLinkedMap(t *testing.T) {
	lm := NewLinkedMap()

	lm.AddLast(node1)
	lm.AddLast(node1)
	lm.RemoveHead()
}
