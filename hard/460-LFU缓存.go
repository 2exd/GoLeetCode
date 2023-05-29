package main

// 双向链表
type doublyLinkedList struct {
	head *doublyListNode
	tail *doublyListNode
}

// 删除尾结点
func (dl *doublyLinkedList) removeTail() {
	pre := dl.tail.prev.prev
	pre.next = dl.tail
	dl.tail.prev = pre
}

// 链表是否为空
func (dl *doublyLinkedList) isEmpty() bool {
	return dl.head.next == dl.tail
}

// 双向链表节点
type doublyListNode struct {
	key       int
	value     int
	frequency int // 使用次数
	prev      *doublyListNode
	next      *doublyListNode
}

// 在某一个节点之前插入一个节点
func addPrev(currentNode *doublyListNode, newNode *doublyListNode) {
	prev := currentNode.prev
	prev.next = newNode
	newNode.next = currentNode
	currentNode.prev = newNode
	newNode.prev = prev
}

// LFUCache 具体的缓存
type LFUCache struct {
	recent   map[int]*doublyListNode // frequency 到使用次数为 frequency 的节点中，最近使用的一个的映射
	count    map[int]int             // frequency 到对应频率的节点数量的映射
	cache    map[int]*doublyListNode // key到节点的映射
	list     *doublyLinkedList       // 双向链表，维护缓存的使用次数（优先）和上一次使用时间
	capacity int                     // 容量
}

func removeNode(node *doublyListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// Constructor 构建缓存容器
func Constructor(capacity int) LFUCache {
	return LFUCache{
		recent:   make(map[int]*doublyListNode),
		cache:    make(map[int]*doublyListNode),
		count:    make(map[int]int),
		list:     newDoublyList(),
		capacity: capacity,
	}
}

func newDoublyList() *doublyLinkedList {
	// head := &doublyListNode{key: 0, value: 0, frequency: 0}
	// tail := &doublyListNode{key: 0, value: 0, frequency: 0}
	head := &doublyListNode{}
	tail := &doublyListNode{}
	head.next = tail
	tail.prev = head
	return &doublyLinkedList{
		head: head,
		tail: tail,
	}
}

// Put 新建缓存
func (lfu *LFUCache) Put(key int, value int) {
	if lfu.capacity == 0 {
		return
	}

	node, ok := lfu.cache[key]
	if ok {
		// key已存在
		lfu.Get(key)
		node.value = value
		return
	}

	// key does not exist
	if len(lfu.cache) >= lfu.capacity {
		// 缓存已满，删除最后一个节点，相应更新cache、count、recent（条件）
		tail := lfu.list.tail.prev
		// update link list
		lfu.list.removeTail()
		// update recent map
		if lfu.count[tail.frequency] <= 1 {
			// 不存在其他freq = n的节点，recent置空
			lfu.recent[tail.frequency] = nil
		}
		// update count map
		lfu.count[tail.frequency]--
		// update cache map
		delete(lfu.cache, tail.key)
	}

	newNode := &doublyListNode{
		key:       key,
		value:     value,
		frequency: 1,
	}

	// insert into linked list
	if lfu.count[1] > 0 {
		addPrev(lfu.recent[1], newNode)
	} else {
		addPrev(lfu.list.tail, newNode)
	}
	// update recent map
	lfu.recent[1] = newNode

	// update count map
	lfu.count[1]++

	// update cache map
	lfu.cache[key] = newNode
}

func (lfu *LFUCache) Get(key int) int {
	if lfu.capacity == 0 {
		return -1
	}

	node, ok := lfu.cache[key]
	if !ok {
		// key doesn't exist in cache
		return -1
	}

	// key exists in cache
	// update link list
	next := node.next
	if lfu.count[node.frequency+1] > 0 {
		// 存在其他使用次数为n+1的缓存，将指定缓存移动到所有使用次数为n+1的节点之前
		removeNode(node)
		addPrev(lfu.recent[node.frequency+1], node)
	} else if lfu.count[node.frequency] > 1 && lfu.recent[node.frequency] != node {
		// 不存在其他使用次数为n+1的缓存，但存在其他使用次数为n的缓存，且当前节点不是最近的节点
		// 将指定缓存移动到所有使用次数为n的节点之前
		removeNode(node)
		addPrev(lfu.recent[node.frequency], node)
	}

	// update recent map
	lfu.recent[node.frequency+1] = node
	if lfu.count[node.frequency] <= 1 {
		// 不存在其他freq = n的节点，recent置空
		lfu.recent[node.frequency] = nil
	} else if lfu.recent[node.frequency] == node {
		// 存在其他freq = n的节点，且recent = node，将recent向后移动一位
		lfu.recent[node.frequency] = next
	}

	// update count map
	lfu.count[node.frequency+1]++
	lfu.count[node.frequency]--

	// update node frequency
	node.frequency++
	return node.value
}

// func main() {
// 	lfu := Constructor(10)
// 	lfu.Put(10, 13)
// 	lfu.Put(3, 17)
// 	lfu.Put(6, 11)
// 	lfu.Put(10, 5)
// 	lfu.Put(9, 10)
// 	fmt.Println(lfu.Get(13))
// 	lfu.Put(2, 19)
// 	fmt.Println(lfu.Get(2))
// 	fmt.Println(lfu.Get(3))
// 	lfu.Put(5, LFUCache)
// 	fmt.Println(lfu.Get(8))
// 	lfu.Put(9, 22)
// 	lfu.Put(5, 5)
// 	lfu.Put(1, 30)
// 	fmt.Println(lfu.Get(11))
// }
