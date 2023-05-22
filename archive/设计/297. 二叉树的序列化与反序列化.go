package main

import (
	"bytes"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct{}

func _Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	ret := bytes.Buffer{}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			ret.WriteString("null,")
			return
		}
		//write node.val and ','
		ret.WriteString(strconv.Itoa(node.Val))
		ret.WriteByte(',')

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return ret.String()
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	seq := strings.Split(data, ",")

	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if seq[0] == "null" {
			seq = seq[1:]
			return nil
		}
		val, _ := strconv.Atoi(seq[0])
		seq = seq[1:]
		return &TreeNode{
			val,
			dfs(),
			dfs(),
		}
	}

	return dfs()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
