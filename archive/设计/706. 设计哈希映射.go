package main

import "container/list"

const Base = 1007

type MyHashMap struct {
	data []list.List
}

type node struct {
	key, value int
}

func nHash(key int) int {
	return key % Base
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{make([]list.List, Base)}
}

/** value will always be non-negative. */
func (hMap *MyHashMap) Put(key int, value int) {
	pos := nHash(key)
	for p := hMap.data[pos].Front(); p != nil; p = p.Next() {
		if temp := p.Value.(node); temp.key == key {
			p.Value = node{key, value}
			return
		}
	}
	hMap.data[pos].PushBack(node{key, value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (hMap *MyHashMap) Get(key int) int {
	pos := nHash(key)
	for p := hMap.data[pos].Front(); p != nil; p = p.Next() {
		if temp := p.Value.(node); temp.key == key {
			return temp.value
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (hMap *MyHashMap) Remove(key int) {
	pos := nHash(key)
	for p := hMap.data[pos].Front(); p != nil; p = p.Next() {
		if temp := p.Value.(node); temp.key == key {
			hMap.data[pos].Remove(p)
			break
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
