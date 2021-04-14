package main

import (
	"./Graph"
	"./Tree"
	"fmt"
)

func main() {
	tree := Tree.NewTree()
	tree.Insert(6)
	tree.Insert(11)
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(8)
	tree.Insert(9)
	tree.Insert(12)
	tree.Insert(21)
	tree.Insert(14)
	tree.Insert(10)
	tree.Insert(19)
	tree.Insert(28)
	tree.Insert(3)
	tree.Insert(17)
	tree.Insert(32)
	tree.Insert(15)
	tree.Insert(16)
	tree.Insert(26)
	tree.Insert(27)
	Graph.GenerateGraph_BTree(tree.Root)
	fmt.Print(":)")
}
