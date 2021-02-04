package Estructura

import (
	"../Persona"
	"fmt"
)

type Node struct {
	next    *Node
	persona *Persona.Client
}

type List struct {
	first *Node
	last  *Node
	count int
}

func NewNode(persona *Persona.Client) *Node {
	return &Node{nil, persona}
}
func Insert(client *Persona.Client, list *List) {
	var newNode *Node = NewNode(client)
	if list.first == nil {
		list.first = newNode
		list.last = newNode
		list.count++
	} else {
		list.last.next = newNode
		list.last = list.last.next
		list.count++
	}
}

func Print(list *List) {
	aux := list.first
	for aux != nil {
		fmt.Println("id: ", aux.persona.Id)
		fmt.Println("name: ", aux.persona.Name)
		fmt.Println("age: ", aux.persona.Age)
		aux = aux.next
	}
}

func NewList() *List {
	return &List{nil, nil, 0}
}
