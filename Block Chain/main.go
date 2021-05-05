package main

import "./structure"

func main() {
	blockChain := structure.NewList()
	blockChain.Insert("hello")
	blockChain.Insert("hello2")
	blockChain.Insert("hello3")
	blockChain.Insert("my data")
	blockChain.Insert("information")
	blockChain.Print()
}
