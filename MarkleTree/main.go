package main

import (
	"./tree"
)

func main() {
	newTree := tree.NewTree()
	newTree.Insert(1)
	newTree.Insert(2)
	newTree.Insert(3)
	newTree.Insert(4)
	newTree.Insert(5)
	newTree.Insert(6)
	newTree.Insert(7)
	newTree.Insert(8)
	newTree.Insert(9)
	newTree.GenerateGraph()
}
