package leetcode

type DNode struct {
	key, val   int
	prev, next *DNode
}

// DoublyLinkedList
type DList struct {
	head, tail *DNode
	Size       int
}

func NewDoublyLinkedList() *DList {
	l := DList{head: &DNode{}, tail: &DNode{}}
	l.head.next = l.tail
	l.tail.prev = l.head
	return &l
}

func (l *DList) AddLast(x *DNode) {
	x.prev = l.tail.prev
	x.next = l.tail

	l.tail.prev.next = x
	l.tail.prev = x

	l.Size++
}

func (l *DList) Remove(x *DNode) {
	x.prev.next = x.next
	x.next.prev = x.prev

	l.Size--
}

func (l *DList) RemoveHead() *DNode {
	if l.head.next == l.tail {
		panic("Remove Head wrong, nil DList!")
	}

	oldHead := l.head.next
	l.Remove(oldHead)
	return oldHead
}

type LinkedMap struct {
	dList     *DList
	keyToNode map[int]*DNode
}

func NewLinkedMap() *LinkedMap {
	l := NewDoublyLinkedList()
	m := make(map[int]*DNode)
	return &LinkedMap{
		dList:     l,
		keyToNode: m,
	}
}

func (l *LinkedMap) Remove(x *DNode) {
	l.dList.Remove(x)
	delete(l.keyToNode, x.key)
}

func (l *LinkedMap) RemoveHead() *DNode {
	h := l.dList.RemoveHead()
	delete(l.keyToNode, h.key)
	return h
}

func (l *LinkedMap) AddLast(x *DNode) {
	if _, ok := l.keyToNode[x.key]; ok {
		panic("wrong! cant add existed node to linked map")
	}

	l.dList.AddLast(x)
	l.keyToNode[x.key] = x
}
