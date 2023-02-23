package template

import (
	"fmt"
)

type Node struct {
	key   int
	value any
	next  *Node
}

func (n *Node) Get(key int) *Node {
	if n == nil {
		return nil
	}
	if n.key == key {
		return n
	}
	return n.next.Get(key)
}

func (n *Node) Put(key int, value any) (*Node, bool) {
	if n == nil {
		return &Node{key: key, value: value, next: nil}, true
	}
	if n.key == key {
		n.value = value
		return n, false
	}
	var new bool
	n.next, new = n.next.Put(key, value)
	return n, new
}

func (n *Node) Remove(key int) *Node {
	if n == nil {
		return nil
	}
	if n.key == key {
		return n.next
	}
	n.next = n.next.Remove(key)
	return n
}

type HashTable struct {
	elmts []*Node
	size  int
}

func NewHashTable(size int) HashTable {
	return HashTable{
		elmts: make([]*Node, size),
		size:  0,
	}
}

func (h HashTable) Hash(key int) int {
	hash := key % len(h.elmts)
	fmt.Println("key: ", key, " / Hash: ", hash)
	return hash
}

func (h *HashTable) Get(key int) any {
	index := h.Hash(key)
	node := h.elmts[index].Get(key)
	if node == nil {
		return nil
	}
	return node.value
}

func (h *HashTable) Put(key int, value any) {
	index := h.Hash(key)
	var new bool
	h.elmts[index], new = h.elmts[index].Put(key, value)
	if new {
		h.size++
	}
	if h.size*2 > len(h.elmts) {
		h.expand()
	}
}

func (h *HashTable) Remove(key int) {
	node := h.Get(key)
	if node == nil {
		return
	}
	index := h.Hash(key)
	h.elmts[index] = h.elmts[index].Remove(key)
	h.size--
}

func (h *HashTable) Size() int {
	return h.size
}

func (h *HashTable) expand() {
	elmts := h.elmts
	h.elmts = make([]*Node, len(elmts)*2)
	h.size = 0
	for _, e := range elmts {
		for n := e; n != nil; n = n.next {
			h.Put(n.key, n.value)
		}
	}
}
