package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type BSTIterator struct {
	vals []int
	ptr  int
}

func Constructor_(root *TreeNode) BSTIterator {
	var it BSTIterator
	it.vals = make([]int, 0)
	it.ptr = 0
	var dfs func(treeNode *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		it.vals = append(it.vals, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return it
}

func (b *BSTIterator) Next() int {
	ret := b.vals[b.ptr]
	b.ptr++
	return ret
}

func (b *BSTIterator) HasNext() bool {
	return b.ptr < len(b.vals)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
