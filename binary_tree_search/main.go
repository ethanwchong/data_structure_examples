package main

import "fmt"

//define a node struct

type Node struct {
	key   int
	left  *Node
	right *Node
}

// create a insert function
func (node *Node) insert(num int) {
	if node != nil && num > node.key {
		if node.right == nil {
			node.right = &Node{key: num}
		} else {
			node.right.insert(num)
		}
	} else {
		if node.left == nil {
			node.left = &Node{key: num}
		} else {
			node.left.insert(num)
		}
	}

}

// search a number from a binary tree
func (node *Node) search(num int) bool {
	if node == nil {
		return false
	}

	if num > node.key {
		return node.right.search(num)
	} else if num < node.key {
		return node.left.search(num)
	}
	return true
}

func main() {
	rootNode := Node{key: 100}
	rootNode.insert(120)
	rootNode.insert(140)
	rootNode.insert(160)
	rootNode.insert(90)
	rootNode.insert(95)
	rootNode.insert(88)
	rootNode.insert(60)
	rootNode.insert(72)
	rootNode.insert(85)
	rootNode.insert(65)
	fmt.Println(rootNode)
	fmt.Println(rootNode.search(81))

}
