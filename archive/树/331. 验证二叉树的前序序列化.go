package main

import (
	"fmt"
	"strings"
)

func isValidSerialization(preorder string) bool {
	if preorder == "#" {
		return true
	}
	nodes := strings.Split(preorder, ",")
	ingress, outgress := 0, 0
	if nodes[0] == "#" {
		return false
	} else {
		outgress = 2
	}
	for i := 1; i < len(nodes); i++ {
		if nodes[i] == "#" {
			ingress++
		} else {
			ingress++
			outgress += 2
		}
		if i != len(nodes)-1 && ingress >= outgress {
			return false
		}
	}
	return ingress == outgress
}

func main() {
	s := "9,3,4,#,#,1,#,#,2,#,6,#,#"
	fmt.Println(isValidSerialization(s))
}
