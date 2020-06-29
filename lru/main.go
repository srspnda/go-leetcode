package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity  int
	evictList *list.List
	items     map[int]*list.Element
}

type item struct {
	key   int
	value int
}

func NewLRUCache(capacity int) LRUCache {
	c := LRUCache{
		capacity:  capacity,
		evictList: list.New(),
		items:     make(map[int]*list.Element),
	}
	return c
}

// Add key into LRU cache, return true if item was evicted
func (c *LRUCache) Put(key, value int) bool {
	if e, ok := c.items[key]; ok {
		c.evictList.MoveToFront(e)
		e.Value.(*item).value = value
		return false
	}

	i := &item{key, value}
	e := c.evictList.PushFront(i)
	c.items[key] = e

	if c.evictList.Len() > c.capacity {
		c.removeOldest()
		return true
	}
	return false
}

// Get key from LRU cache, if key exists move to front of eviction list and
// return the value of the item key, else return -1 if key does not exist
func (c *LRUCache) Get(key int) int {
	if e, ok := c.items[key]; ok {
		c.evictList.MoveToFront(e)
		return e.Value.(*item).value
	}
	return -1
}

// Return a slice of keys ordered from oldest to youngest in the LRU cache
func (c *LRUCache) Keys() []int {
	keys := make([]int, len(c.items))
	i := 0
	for e := c.evictList.Back(); e != nil; e = e.Prev() {
		keys[i] = e.Value.(*item).key
		i++
	}
	return keys
}

// Remove oldest element from the eviction list
func (c *LRUCache) removeOldest() {
	e := c.evictList.Back()
	if e != nil {
		c.removeElement(e)
	}
}

// Remove element from eviction list, and delete key out of LRU cache items
func (c *LRUCache) removeElement(e *list.Element) {
	c.evictList.Remove(e)
	i := e.Value.(*item)
	delete(c.items, i.key)
}

func main() {
	l := NewLRUCache(2)
	l.Put(1, 1)
	l.Put(2, 2)
	fmt.Println(l.Keys())
	fmt.Println(l.Get(1))
	fmt.Println(l.Keys())
	l.Put(3, 3)
	fmt.Println(l.Keys())
	fmt.Println(l.Get(2))
	fmt.Println(l.Keys())
	l.Put(4, 4)
	fmt.Println(l.Keys())
	fmt.Println(l.Get(1))
	fmt.Println(l.Keys())
	fmt.Println(l.Get(3))
	fmt.Println(l.Keys())
	fmt.Println(l.Get(4))
	fmt.Println(l.Keys())
}
