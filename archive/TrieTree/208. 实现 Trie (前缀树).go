package TrieTree

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

// Constructor /** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

// Insert /** Inserts a word into the trie. */
func (trie *Trie) Insert(word string) {
	node := trie
	for _, ch := range word {
		index := ch - 'a'
		if node.children[index] == nil {
			node.children[index] = new(Trie)
		}
		node = node.children[index]
	}
	node.isEnd = true
}

// Search /** Returns if the word is in the trie. */
func (trie *Trie) Search(word string) bool {
	node := trie
	for _, ch := range word {
		index := ch - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	if node.isEnd == false {
		return false
	}
	return true
}

// StartsWith /** Returns if there is any word in the trie that starts with the given prefix. */
func (trie *Trie) StartsWith(prefix string) bool {
	node := trie
	for _, ch := range prefix {
		index := ch - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
