/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start
type LRUCache struct {
	size, c int
	cache map[int]*DLinkedList
	head,tail *DLinkedList
}

type DLinkedList struct {
	key, val int
	next,prev *DLinkedList
}

func initDLinkedList(key,val int) *DLinkedList {
	return &DLinkedList {
		key:key,
		val:val,
	}

}


func Constructor(capacity int) LRUCache {
	cache := LRUCache {
		c : capacity,
		cache: make(map[int]*DLinkedList),
		head : initDLinkedList(0,0),
		tail: initDLinkedList(0,0),
	}

	cache.head.next= cache.tail
	cache.tail.prev = cache.head

	return cache
}


func (this *LRUCache) Get(key int) int {
	n,ok := this.cache[key]
	if !ok {
		return -1
	}

	this.moveToHead(n)
	return n.val
}


func (this *LRUCache) Put(key int, value int)  {
	n,ok := this.cache[key]
	if ok {
		n.val=value
		this.moveToHead(n)
		return
	}

	n = initDLinkedList(key,value)
	this.cache[key]=n
	this.addToHead(n)
	this.size++

	if this.size > this.c {
		temp := this.removeTail()
		delete(this.cache,temp.key)
		this.size--
	}
}

func (this *LRUCache) addToHead(n *DLinkedList) {
	n.prev = this.head
	n.next = this.head.next
	this.head.next.prev = n
	this.head.next=n
}

func (this *LRUCache) removeNode(n *DLinkedList) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (this *LRUCache) moveToHead(n *DLinkedList) {
	this.removeNode(n)
	this.addToHead(n)
}

func (this *LRUCache) removeTail() *DLinkedList {
	n := this.tail.prev
	this.removeNode(n)

	return n
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

