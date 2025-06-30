package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type entry struct {
	key   int
	value int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (lru *LRUCache) Get(key int) int {
	if elem, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(elem)
		return elem.Value.(*entry).value
	}
	return -1
}

func (lru *LRUCache) Insert(key int, value int) {
	if elem, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}

	if lru.list.Len() >= lru.capacity {
		back := lru.list.Back()
		delete(lru.cache, back.Value.(*entry).key)
		lru.list.Remove(back)
	}

	elem := lru.list.PushFront(&entry{key, value})
	lru.cache[key] = elem
}

func main() {

	cache := NewLRUCache(2)

	cache.Insert(1, 100)
	cache.Insert(2, 200)
	fmt.Println(cache.Get(1))
	cache.Insert(3, 300)
	fmt.Println(cache.Get(2))
}
