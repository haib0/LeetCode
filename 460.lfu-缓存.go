/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU 缓存
 */
package leetcode

// @lc code=start
import (
	"fmt"
	"log"
)

type LFUCache struct {
	keyToVal  map[int]int
	keyToFreq map[int]int
	freqToLM  map[int]*LinkedMap
	minFreq   int
	cap       int
}

func Constructor460(capacity int) LFUCache {
	f := make(map[int]*LinkedMap)
	f[1] = NewLinkedMap()
	return LFUCache{
		keyToVal:  make(map[int]int),
		keyToFreq: make(map[int]int),
		freqToLM:  f,
		minFreq:   1,
		cap:       capacity,
	}
}

func (this *LFUCache) increaseFreq(key int) {
	freq, ok := this.keyToFreq[key]
	if !ok {
		log.Panicf("Wrong freq, key=%d, keyToFreq=%v", key, this.keyToFreq)
	}

	lM, ok := this.freqToLM[freq]
	fmt.Println(lM)
	if !ok {
		log.Panicf("Wrong Lnked Map, freq=%d", freq)
	}

	newLM, ok := this.freqToLM[freq+1]
	if !ok {
		newLM = NewLinkedMap()
		this.freqToLM[freq+1] = newLM
	}

	node, ok := lM.keyToNode[key]
	if !ok {
		log.Panicf("Wrong node, key=%d", key)
	}

	newLM.AddLast(node)
	lM.Remove(node)
	this.keyToFreq[key] = freq + 1

	if len(lM.keyToNode) == 0 {
		// delete(this.freqToLM, freq)
		if freq == this.minFreq {
			this.minFreq = freq + 1
		}
	}
}

func (this *LFUCache) Get(key int) int {
	v, ok := this.keyToVal[key]
	if !ok {
		return -1
	}
	this.increaseFreq(key)
	return v
}

func (this *LFUCache) Put(key int, value int) {
	_, ok := this.keyToVal[key]
	if ok {
		this.keyToVal[key] = value
		this.increaseFreq(key)
	} else {
		if len(this.keyToVal) >= this.cap {
			minLM := this.freqToLM[this.minFreq]
			h := minLM.RemoveHead()
			// if len(minLM.keyToNode) == 0 {
			// 	delete(this.freqToLM, this.minFreq)
			// }
			delete(this.keyToVal, h.key)
			delete(this.keyToFreq, h.key)
		}
		
		this.keyToVal[key] = value
		this.keyToFreq[key] = 1
		lm1 := this.freqToLM[1]
		lm1.AddLast(&DNode{
			key: key, val: value,
		})
		this.minFreq = 1
	}

}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
