package main

import (
	"fmt"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	type trieNode struct {
		char     rune
		isEnding bool
		children map[rune]*trieNode
	}
	newNode := func(ch rune) *trieNode {
		return &trieNode{
			char:     ch,
			isEnding: false,
			children: make(map[rune]*trieNode),
		}
	}

	root := newNode('/')
	for _, word := range dictionary {
		node := root
		for _, ch := range word {
			val, ok := node.children[ch]
			if !ok {
				val = newNode(ch)
				node.children[ch] = val
			}
			node = val
		}
		node.isEnding = true
	}

	pieces := strings.Split(sentence, " ")
	for i, word := range pieces {
		node := root
		for j, ch := range word {
			if val, ok := node.children[ch]; ok {
				node = val
				if val.isEnding {
					pieces[i] = word[:j+1]
					break
				}
			} else {
				break
			}
		}
	}
	return strings.Join(pieces, " ")
}

func main() {
	dict := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	if replaceWords(dict, sentence) == "the cat was rat by the bat" {
		fmt.Println("passed test case")
	} else {
		fmt.Println("failed to pass test case ")
	}
}
