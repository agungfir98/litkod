package main

type Node struct {
	prev *Node
	next *Node
	key  int
	val  int
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	tail     *Node
	head     *Node
}

func Constructor(capacity int) LRUCache {
	tail := &Node{}
	head := &Node{}
	head.next = tail
	tail.prev = head
	cache := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		tail:     tail,
		head:     head,
	}

	return cache
}

func (this *LRUCache) remove(node *Node) {
	prev, next := node.prev, node.next

	prev.next = next
	next.prev = prev
}

func (this *LRUCache) addToHead(node *Node) {
	node.prev = this.head
	node.next = this.head.next

	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *Node) {
	this.remove(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *Node {
	lru := this.tail.prev
	this.remove(lru)
	return lru
}

func (this *LRUCache) Get(key int) int {
	cache, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.moveToHead(cache)
	return cache.val
}

func (this *LRUCache) Put(key int, value int) {
	cache, ok := this.cache[key]
	if ok {
		cache.val = value
		this.moveToHead(cache)
	} else {

		node := &Node{
			key: key,
			val: value,
		}
		this.cache[key] = node
		this.addToHead(node)

		if len(this.cache) > this.capacity {
			lru := this.removeTail()
			delete(this.cache, lru.key)
		}
	}

}
