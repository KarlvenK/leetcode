package search_hard

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	var ans string
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			ans += "null/"
			return
		}
		ans += strconv.Itoa(node.Val)
		ans += "/"
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ans[:len(ans)-1]
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	slc := strings.Split(data, "/")
	var build func() *TreeNode
	build = func() *TreeNode {
		if slc[0] == "null" {
			slc = slc[1:]
			return nil
		}
		val, _ := strconv.Atoi(slc[0])
		slc = slc[1:]
		return &TreeNode{
			val,
			build(),
			build(),
		}
	}
	return build()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

/*
前序遍历将它序列化
然后逆序列化的时候，模拟逆向前序遍历
*/
