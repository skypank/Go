package main

import "fmt"

type Tree struct {
	data  int
	left  *Tree
	right *Tree
}

func main() {
	// declare a tree
	var rootNode *Tree
	// insert first node in tree
	rootNode = &Tree{data: 3, left: nil, right: nil}
	// create tree with few elements
	element := []int{2, 1, 4, 5}
	for _, num := range element {
		insert(rootNode, num)
	}
	// print the tree
	//printTree(rootNode, 0)

	// traversal techniques
	// 1. pre order traversal // root,left,right
	fmt.Println(" pre order traversal")
	preOrderTraversal(rootNode)
	// 2. post order traversal // left,right,root
	fmt.Println(" post order traversal")
	postOrderTraversal(rootNode)
	// 3. in order traversal // left,root,right
	fmt.Println(" in order traversal")
	inOrderTraversal(rootNode)
	// 4. level order traversal // level by level
	fmt.Println(" level order traversal")
	fmt.Println(rootNode.data)
	levelOrderTraversal(rootNode)

}

// insert on a root of tree
func insert(root *Tree, num int) {
	if root == nil {
		root = &Tree{data: num, left: nil, right: nil}
	} else if num < root.data {
		if root.left == nil {
			root.left = &Tree{data: num, left: nil, right: nil}
		} else {
			insert(root.left, num)
		}
	} else {
		if root.right == nil {
			root.right = &Tree{data: num, left: nil, right: nil}
		} else {
			insert(root.right, num)
		}
	}
}

// print tree formatting
func printTree(n *Tree, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "      "
		}
		format += "-<"
		level++
		printTree(n.left, level)
		fmt.Printf(format+"%d\n", n.data)
		printTree(n.right, level)
	}
}

func preOrderTraversal(root *Tree) {
	if root == nil {
		return
	}
	fmt.Println(root.data)
	preOrderTraversal(root.left)
	preOrderTraversal(root.right)
}

func postOrderTraversal(root *Tree) {
	if root == nil {
		return
	}
	postOrderTraversal(root.left)
	postOrderTraversal(root.right)
	fmt.Println(root.data)
}

func inOrderTraversal(root *Tree) {
	if root == nil {
		return
	}
	inOrderTraversal(root.left)
	fmt.Println(root.data)
	inOrderTraversal(root.right)
}

var queue []int

func levelOrderTraversal(root *Tree) {
	if root == nil {
		return
	}
	if root.left != nil {
		fmt.Println(root.left.data)
	}
	if root.right != nil {
		fmt.Println(root.right.data)
	}
	levelOrderTraversal(root.left)
	levelOrderTraversal(root.right)
}
