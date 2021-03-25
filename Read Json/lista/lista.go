package lista

import (
	"fmt"
	"../tienda"
)

type Node struct {
	store *tienda.Store
	Next   *Node
	Back   *Node
}

type List struct {
	first, last *Node
	count int
}
func NewList() *List {
	return &List{nil, nil,0}
}
func NewNode(store *tienda.Store) *Node {
	return &Node{store, nil,nil}
}
func Insert(store *tienda.Store, list *List) {
	var newNode = NewNode(store)
	if list.first == nil {
		list.first = newNode
		list.last = newNode
		list.count++
	} else {
		list.last.Next = newNode
		newNode.Back = list.last
		list.last = newNode
		list.count++
	}
}

func Print(list *List) {
	aux := list.first
	for aux != nil {
		fmt.Print("[", aux.store.Name, " - ", aux.store.Id, " - ", aux.store.Age, "] -->")
		aux = aux.Next
	}
	fmt.Println()
}
