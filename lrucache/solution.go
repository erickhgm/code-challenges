package template

import "fmt"

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	items    map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		items:    map[int]*Node{},
	}
}

func (c *LRUCache) Get(key int) int {
	node, ok := c.items[key]
	if !ok {
		return -1
	}
	if node == c.head {
		return node.value
	}
	if node == c.tail {
		c.tail = node.prev
	}

	nPrev := node.prev
	nNext := node.next
	if nPrev != nil {
		nPrev.next = node.next
	}
	if nNext != nil {
		nNext.prev = node.prev
	}

	head := c.head
	head.prev = node
	node.next = head
	node.prev = nil

	c.head = node

	return node.value
}

func (c *LRUCache) Put(key int, value int) {
	node, ok := c.items[key]
	if ok {
		node.value = value

	} else {
		node = &Node{key: key, value: value}

		if len(c.items) >= c.capacity {
			delete(c.items, c.tail.key)
			tail := c.tail
			c.tail = tail.prev
			if c.tail != nil {
				c.tail.next = nil
			}
		}
		if c.head == nil {
			c.head = node
			c.tail = node
		} else {
			node.next = c.head
			c.head.prev = node
			c.head = node
		}
		c.items[key] = node
	}
}

func (c *LRUCache) Print() {
	stopValue := c.tail.value
	node := c.head
	for node.value != stopValue {
		fmt.Println(node.value)
		node = node.next
		stopValue = node.value
	}
	fmt.Println(node.value)
}
