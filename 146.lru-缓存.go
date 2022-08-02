/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */
package leetcode

// @lc code=start

// import "log"

// type DNode struct {
// 	key, val   int
// 	prev, next *DNode
// }

// // add node x before p
// func (p *DNode) AddBefore(x *DNode) {
// 	if p == nil || x == nil {
// 		log.Panic("add before error!")
// 	}

// 	x.prev = p.prev
// 	x.next = p
// 	p.prev.next = x
// 	p.prev = x
// }

// // add node x after p
// func (p *DNode) AddAfter(x *DNode) {
// 	if p == nil || x == nil {
// 		log.Panic("add after error!")
// 	}

// 	x.prev = p
// 	x.next = p.next
// 	p.next.prev = x
// 	p.next = x
// }

// func (p *DNode) Remove() {
// 	if p == nil || p.prev == nil || p.next == nil {
// 		log.Panic("remove error!")
// 	}
// 	p.prev.next = p.next
// 	p.next.prev = p.prev
// }

// // DoublyLinkedList
// type DList struct {
// 	head, tail *DNode
// 	Size       int
// }

// func NewDoublyLinkedList() *DList {
// 	l := DList{head: &DNode{}, tail: &DNode{}}
// 	l.head.next = l.tail
// 	l.tail.prev = l.head
// 	return &l
// }

// func (l *DList) AddLast(x *DNode) {
// 	l.tail.AddBefore(x)
// 	l.Size++
// }

// func (l *DList) Remove(x *DNode) {
// 	x.Remove()
// 	l.Size--
// }

// func (l *DList) RemoveHead() *DNode {
// 	if l.head.next == l.tail {
// 		return nil
// 	}

// 	oldHead := l.head.next
// 	l.Remove(oldHead)
// 	return oldHead
// }

type LRUCache struct {
	cache    *DList
	cacheMap map[int]*DNode
	cap      int
}

func Constructor146(capacity int) LRUCache {
	cache := NewDoublyLinkedList()
	cacheMap := make(map[int]*DNode)
	return LRUCache{cache: cache, cacheMap: cacheMap, cap: capacity}
}

func (this *LRUCache) MakeRecentlt(node *DNode) {
	this.cache.Remove(node)
	this.cache.AddLast(node)
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.cacheMap[key]
	if !ok {
		return -1
	}
	this.MakeRecentlt(v)
	return v.val
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.cacheMap[key]
	if !ok {
		if this.cache.Size >= this.cap {
			oldHead := this.cache.RemoveHead()
			delete(this.cacheMap, oldHead.key)
		}
		newNode := &DNode{key: key, val: value}
		this.cache.AddLast(newNode)
		this.cacheMap[key] = newNode
	} else {
		v.val = value
		this.MakeRecentlt(v)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
