package main

import (
	"container/list"
	"fmt"
)

/*
请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。

实现 LFUCache 类：

LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
int get(int key) - 如果键 key 存在于缓存中，则获取键的值，否则返回 -1 。
void put(int key, int value) - 如果键 key 已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量 capacity 时，
则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最近最久未使用 的键。
为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。

当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。

函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

输入：
["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
输出：
[null, null, null, 1, null, -1, 3, null, -1, 3, 4]

解释：
// cnt(x) = 键 x 的使用计数
// cache=[] 将显示最后一次使用的顺序（最左边的元素是最近的）
LFUCache lfu = new LFUCache(2);
lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
lfu.get(1);      // 返回 1
                 // cache=[1,2], cnt(2)=1, cnt(1)=2
lfu.put(3, 3);   // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
                 // cache=[3,1], cnt(3)=1, cnt(1)=2
lfu.get(2);      // 返回 -1（未找到）
lfu.get(3);      // 返回 3
                 // cache=[3,1], cnt(3)=2, cnt(1)=2
lfu.put(4, 4);   // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
                 // cache=[4,3], cnt(4)=1, cnt(3)=2
lfu.get(1);      // 返回 -1（未找到）
lfu.get(3);      // 返回 3
                 // cache=[3,4], cnt(4)=1, cnt(3)=3
lfu.get(4);      // 返回 4
                 // cache=[3,4], cnt(4)=2, cnt(3)=3
*/

// 用了切片，不太完善
//type LFUCache struct {
//	keyToElement map[int]*list.Element
//	keyToFreq    map[int]int
//	FreqToList   map[int]*list.List
//	minFreq      int
//	cap          int
//}
//
//func Constructor(capacity int) LFUCache {
//	return LFUCache{
//		keyToElement: make(map[int]*list.Element),
//		keyToFreq:    make(map[int]int),
//		FreqToList:   make(map[int]*list.List),
//		minFreq:      0,
//		cap:          capacity,
//	}
//}
//
//func (this *LFUCache) Get(key int) int {
//	if element, ok := this.keyToElement[key]; ok {
//		this.increaseFreq(key)
//		val := element.Value.([]int)[1]
//		return val
//	} else {
//		return -1
//	}
//}
//
//func (this *LFUCache) Put(key int, value int) {
//	if this.cap <= 0 {
//		return
//	}
//	if _, ok := this.keyToElement[key]; ok {
//		this.keyToElement[key].Value = []int{key, value}
//		this.increaseFreq(key)
//		return
//	}
//
//	if this.cap <= len(this.keyToElement) {
//		this.removeMinFreq()
//	}
//
//	if _, ok := this.FreqToList[1]; !ok {
//		this.FreqToList[1] = &list.List{}
//	}
//
//	newElement := this.FreqToList[1].PushBack([]int{key, value})
//	this.keyToElement[key] = newElement
//	this.keyToFreq[key] = 1
//	this.minFreq = 1
//	return
//
//}
//
//func (this *LFUCache) increaseFreq(key int) {
//	freq := this.keyToFreq[key]
//	this.keyToFreq[key]++
//
//	this.FreqToList[freq].Remove(this.keyToElement[key])
//	if length := this.FreqToList[freq].Len(); 0 == length {
//		if this.minFreq == freq {
//			this.minFreq++
//		}
//		delete(this.FreqToList, freq)
//	}
//
//	if _, ok := this.FreqToList[freq+1]; !ok {
//		this.FreqToList[freq+1] = &list.List{}
//	}
//	newElement := this.FreqToList[freq+1].PushBack(this.keyToElement[key].Value.([]int))
//	this.keyToElement[key] = newElement
//}
//
//func (this *LFUCache) removeMinFreq() {
//	deletingElement := this.FreqToList[this.minFreq].Front()
//
//	val := this.FreqToList[this.minFreq].Remove(deletingElement)
//	if length := this.FreqToList[this.minFreq].Len(); 0 == length {
//		delete(this.FreqToList, this.minFreq)
//	}
//
//	delete(this.keyToElement, val.([]int)[0])
//	delete(this.keyToFreq, val.([]int)[0])
//}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type node struct {
	key   int
	freq  int
	value int
}

type LFUCache struct {
	keyToElement map[int]*list.Element
	FreqToList   map[int]*list.List // freq对应一条时间链
	minFreq      int
	cap          int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		keyToElement: make(map[int]*list.Element),
		FreqToList:   make(map[int]*list.List),
		minFreq:      0,
		cap:          capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if element, ok := this.keyToElement[key]; ok {
		this.increaseFreq(key)
		return element.Value.(*node).value
	} else {
		return -1
	}
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap <= 0 {
		return
	}

	// 先判断是否存在
	if element, ok := this.keyToElement[key]; ok {
		// 下面先后无所谓,因为Value指向的node地址相同
		element.Value.(*node).value = value
		this.increaseFreq(key)
		return
	}

	if this.cap <= len(this.keyToElement) {
		this.removeMinFreq()
	}

	newNode := &node{
		key:   key,
		freq:  1,
		value: value,
	}

	this.createEmptyList(1)
	newElement := this.FreqToList[1].PushBack(newNode)
	this.keyToElement[key] = newElement
	this.minFreq = 1
	return
}

func (this *LFUCache) increaseFreq(key int) {
	changingElement := this.keyToElement[key]
	freq := changingElement.Value.(*node).freq

	this.FreqToList[freq].Remove(changingElement)
	this.deleteEmptyList(freq)

	this.createEmptyList(freq + 1)
	newElement := this.FreqToList[freq+1].PushBack(changingElement.Value.(*node))
	this.keyToElement[key] = newElement
	newElement.Value.(*node).freq++
}

func (this *LFUCache) removeMinFreq() {
	deletingElement := this.FreqToList[this.minFreq].Front()
	this.FreqToList[this.minFreq].Remove(deletingElement)
	this.deleteEmptyList(this.minFreq)

	deletingElementKey := deletingElement.Value.(*node).key
	delete(this.keyToElement, deletingElementKey)
}

func (this *LFUCache) deleteEmptyList(freq int) {
	if this.FreqToList[freq].Len() == 0 {
		delete(this.FreqToList, freq)
		if this.minFreq == freq {
			this.minFreq++
		}
	}
}

func (this *LFUCache) createEmptyList(freq int) {
	if _, ok := this.FreqToList[freq]; !ok {
		this.FreqToList[freq] = &list.List{}
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	//if a:=1;a==1 {
	//
	//}
	//fmt.Println(a)
	//a := &list.Element{Value: 1}
	//b := &list.Element{Value: 2}
	//c := &list.Element{Value: 3}
	//d := &list.Element{Value: 4}
	//test1 := &list.List{}
	//test1.PushBack(a)
	//test1.PushBack(b)
	//test1.PushBack(c)
	//test1.PushBack(d)
	//test2 := &list.List{}
	//
	//test1.Remove(c)
	//test2.PushBack(c)
	//// pushback后b value仍存在，但是next和prev节点不存在
	////fmt.Println(b.Next())
	////fmt.Println(b.Prev())
	//
	//test1node := test1.Front()
	//
	//test2node := test2.Front()
	//for nil != test1node {
	//	fmt.Println(test1node.Value.(*list.Element).Value)
	//	test1node = test1node.Next()
	//}
	//
	//for nil != test2node {
	//	fmt.Println(test2node.Value)
	//	test2node = test2node.Next()
	//}

	keyToElement := make(map[int]*list.Element)
	a := 1
	b := 2
	c := 3
	d := 4
	test1 := &list.List{}
	test1.PushBack(a)
	test1.PushBack(b)
	a1 := test1.PushBack(a)
	keyToElement[a] = a1

	test1.PushBack(d)
	test2 := &list.List{}

	a1 = test1.Front()
	a2 := test1.Remove(a1)

	test2.PushBack(c)

	fmt.Println(test1.Len())
	fmt.Println(keyToElement)
	delete(keyToElement, a2.(int))
	val, ok := keyToElement[a2.(int)]
	if ok {
	} else {
		fmt.Println("aaaaaaaaaaaaaaaaaaaaaaa")
	}
	fmt.Println(val)
	fmt.Println(ok)
	fmt.Println("==============")

	fmt.Println(keyToElement)
	fmt.Println(a2.(int))
	// pushback后b value仍存在，但是next和prev节点不存在
	//fmt.Println(b.Next())
	//fmt.Println(b.Prev())

	//test1node := test1.Front()
	//
	//test2node := test2.Front()
	//for nil != test1node {
	//	fmt.Println(test1node.Value.(*list.Element).Value)
	//	test1node = test1node.Next()
	//}
	//
	//for nil != test2node {
	//	fmt.Println(test2node.Value)
	//	test2node = test2node.Next()
	//}
	//
	//
	//test11111 := &list.List{}
	//test11111.PushBack(1,1)
}
