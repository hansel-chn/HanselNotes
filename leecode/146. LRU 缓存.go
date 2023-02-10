package main

import (
	"container/list"
)

/*
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*/

type LRUNode struct {
	key int
	val int
}

type LRUCache struct {
	keyToElement map[int]*list.Element
	seq          *list.List
	cap          int
}

func LRUConstructor(capacity int) LRUCache {
	return LRUCache{
		keyToElement: make(map[int]*list.Element),
		seq:          &list.List{},
		cap:          capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if element, ok := this.keyToElement[key]; ok {
		this.keyToElement[key] = this.refreshSeq(element)
		return this.keyToElement[key].Value.(*LRUNode).val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if this.cap <= 0 {
		return
	}

	if element, ok := this.keyToElement[key]; ok {
		// err has nothing to do with map
		// 断言得到的值不能寻址，这算右值嘛？但是传指针就可以。
		// go里好像就根据两个来处理，一能不能寻址，满足一就可以赋值；二是要注意值拷贝，有时可能没有赋值给想要的地方
		// element.Value.(LRUNode).val = value

		element.Value.(*LRUNode).val = value
		this.keyToElement[key] = this.refreshSeq(element)
		return
	}

	if this.cap <= len(this.keyToElement) {
		this.removeFromSeq()
	}

	newElement := this.seq.PushBack(&LRUNode{key: key, val: value})
	this.keyToElement[key] = newElement
}

func (this *LRUCache) refreshSeq(element *list.Element) *list.Element {
	this.seq.Remove(element)
	return this.seq.PushBack(element.Value.(*LRUNode))
}

func (this *LRUCache) removeFromSeq() {
	removingElement := this.seq.Front()
	this.seq.Remove(removingElement)

	delete(this.keyToElement, removingElement.Value.(*LRUNode).key)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {

}
