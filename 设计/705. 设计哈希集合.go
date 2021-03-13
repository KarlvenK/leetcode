package main

import (
	"container/list"
)

const base = 1007

type MyHashSet struct {
	data []list.List
}

func hash(key int) int {
	return key % base
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{make([]list.List, base)}
}

func (set *MyHashSet) Add(key int) {
	if !set.Contains(key) {
		pos := hash(key)
		set.data[pos].PushBack(key)
	}
}

func (set *MyHashSet) Remove(key int) {
	pos := hash(key)
	for p := set.data[pos].Front(); p != nil; p = p.Next() {
		if p.Value.(int) == key {
			set.data[pos].Remove(p)
			break
		}
	}
}

/** Returns true if this set contains the specified element */
func (set *MyHashSet) Contains(key int) bool {
	pos := hash(key)
	for p := set.data[pos].Front(); p != nil; p = p.Next() {
		if p.Value.(int) == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
