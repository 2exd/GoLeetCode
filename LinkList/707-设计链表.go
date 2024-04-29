package main

type MyLinkedList struct {
	head, tail *MyLinkListNode
	size       int
}

type MyLinkListNode struct {
	value      int
	prev, next *MyLinkListNode
}

func Constructor() MyLinkedList {

	l := &MyLinkedList{
		size: 0,
		head: &MyLinkListNode{},
		tail: &MyLinkListNode{},
	}

	l.head.next = l.tail
	l.tail.prev = l.head

	return *l
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.size {
		return -1
	}
	i := 0
	cur := this.head.next
	for cur != this.tail {
		if i == index {
			return cur.value
		}
		cur = cur.next
		i++
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &MyLinkListNode{
		value: val,
	}
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &MyLinkListNode{
		value: val,
	}
	node.next = this.tail
	node.prev = this.tail.prev
	this.tail.prev.next = node
	this.tail.prev = node
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	if index == this.size {
		this.AddAtTail(val)
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	cur := this.head.next
	i := 0
	for i != index {
		cur = cur.next
		i++
	}

	node := &MyLinkListNode{
		value: val,
	}

	node.next = cur
	node.prev = cur.prev
	cur.prev.next = node
	cur.prev = node

	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size {
		return
	}
	i := 0
	cur := this.head.next
	for i != index {
		cur = cur.next
		i++
	}
	cur.prev.next = cur.next
	cur.next.prev = cur.prev
	this.size--
}

// func main() {
// 	obj := Constructor()
// 	obj.AddAtHead(1)
// 	obj.AddAtTail(3)
// 	obj.AddAtIndex(1, 2)
// 	obj.DeleteAtIndex(1)
// 	println(obj.Get(1))
// }
