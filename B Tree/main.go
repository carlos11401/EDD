package main

import (
	"./Tree"
	"fmt"
)

func main() {
	tree := Tree.NewTree()
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(40)
	tree.Insert(50)
	tree.Insert(60)
	tree.Insert(15)
	fmt.Print(":)")
}
