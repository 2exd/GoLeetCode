package main

type LRUCache struct {
	cache      map[int]*DLinkListNode
	head, tail *DLinkListNode
	size       int
	capacity   int
}

type DLinkListNode struct {
	key, value int
	prev, next *DLinkListNode
}

func initDLinkList(value, key int) *DLinkListNode {
	return &DLinkListNode{
		key:   key,
		value: value,
	}
}

func Constructor146(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		size:     0,
		cache:    map[int]*DLinkListNode{},
		head:     initDLinkList(0, 0),
		tail:     initDLinkList(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	// check if the key exists
	if _, ok := this.cache[key]; !ok {
		// not found
		return -1
	}

	// move to head
	node := this.cache[key]
	this.moveHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	// check if the key exists
	if _, ok := this.cache[key]; !ok {
		// init node
		node := initDLinkList(value, key)
		// add to map
		this.cache[key] = node
		// add to head
		this.addNodeToHead(node)
		this.size++

		// check if oversize
		if this.size > this.capacity {
			// remove tail node
			remove := this.removeTail()
			// map remove key value
			delete(this.cache, remove.key)
			this.size--
		}
	} else {
		// move to head
		node := this.cache[key]
		// update
		node.value = value
		this.moveHead(node)
	}
}

func (this *LRUCache) addNodeToHead(node *DLinkListNode) {
	node.prev = this.head
	node.next = this.head.next

	this.head.next.prev = node
	this.head.next = node

}

func (this *LRUCache) removeNode(node *DLinkListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveHead(node *DLinkListNode) {
	// 先删除原有节点
	this.removeNode(node)
	// add to head
	this.addNodeToHead(node)
}

func (this *LRUCache) removeTail() *DLinkListNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

// func main() {
// 	lRUCache := Constructor(2)
// 	lRUCache.Put(1, 1)           // 缓存是 {1=1}
// 	lRUCache.Put(2, 2)           // 缓存是 {1=1, 2=2}
// 	fmt.Println(lRUCache.Get(1)) // 返回 1
// 	lRUCache.Put(3, 3)           // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
// 	fmt.Println(lRUCache.Get(2)) // 返回 -1 (未找到)
// 	lRUCache.Put(4, 4)           // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}   fmt.Println(lRUCache.Get(1)) // 返回 -1 (未找到)
// 	fmt.Println(lRUCache.Get(3)) // 返回 3   fmt.Println(lRUCache.Get(4)) // 返回 4}
//
// }
