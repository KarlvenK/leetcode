package main

import (
	"math"
	"math/rand"
	"time"
)

const (
	maxLevel    int     = 18
	probability float64 = 1 / math.E
)

var probTable []float64

func init() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < maxLevel; i++ {
		prob := math.Pow(probability, float64(i))
		probTable = append(probTable, prob)
	}
}

type Skiplist struct {
	curLevel int
	Head     *SkipNode
}

type SkipNode struct {
	Val  int
	Prev []*SkipNode
	Next []*SkipNode
}

func Constructor__() Skiplist {
	head := &SkipNode{}
	for i := 0; i < maxLevel; i++ {
		head.Next = append(head.Next, nil)
	}
	return Skiplist{
		curLevel: 0,
		Head:     head,
	}
}

func randomLevel() int {
	level := 0
	r := float64(rand.Int63()) / (1 << 63)
	for r < probTable[level] {
		level++
		if level == maxLevel {
			break
		}
	}
	return level
}

func newNode(val int, level int) *SkipNode {
	next := make([]*SkipNode, level+1)
	prev := make([]*SkipNode, level+1)
	node := &SkipNode{
		Val:  val,
		Prev: prev,
		Next: next,
	}
	return node
}

func (skl *Skiplist) search(target int) *SkipNode {
	node := skl.Head
	for i := skl.curLevel; i >= 0; i-- {
		for node.Next[i] != nil && node.Next[i].Val <= target {
			node = node.Next[i]
		}
		if node != skl.Head && node.Val == target {
			return node
		}
	}
	return nil
}

func (skl *Skiplist) Search(target int) bool {
	res := skl.search(target)
	return res != nil
}

func (skl *Skiplist) Add(num int) {
	level := randomLevel()
	node := newNode(num, level)
	updateN := skl.Head
	for i := skl.curLevel; i >= 0; i-- {
		updateN = skl.findClosest(updateN, i, num)
		if i <= level {
			node.Prev[i] = updateN
			node.Next[i] = updateN.Next[i]
			updateN.Next[i] = node
			if node.Next[i] != nil {
				node.Next[i].Prev[i] = node
			}
		}
	}
	if level > skl.curLevel {
		for i := skl.curLevel + 1; i <= level; i++ {
			skl.Head.Next[i] = node
			node.Prev[i] = skl.Head
		}
		skl.curLevel = level
	}
}

func (skl *Skiplist) findClosest(n *SkipNode, level int, val int) *SkipNode {
	for n.Next[level] != nil && n.Next[level].Val < val {
		n = n.Next[level]
	}
	return n
}

func (skl *Skiplist) Erase(num int) bool {
	node := skl.search(num)
	if node == nil {
		return false
	}
	for i := 0; i < len(node.Next); i++ {
		pre, nex := node.Prev[i], node.Next[i]
		pre.Next[i] = nex
		if nex != nil {
			nex.Prev[i] = pre
		}
	}
	return true
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
