package Estructura

import (
	"fmt"

	"../Persona"
)

type Node struct {
	next    *Node
	persona *Persona.Cliente
}

type List struct {
	first *Node
	last  *Node
	count int
}

func New_Nodo(persona *Persona.Cliente) *Node {
	return &Node{nil, persona}
}
func Insert(client *Persona.Cliente, list *List) {
	var new *Node = New_Nodo(client)
	if list.first == nil {
		list.first = new
		list.last = new
		list.count++
	} else {
		list.last.next = new
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

func New_List() *List {
	return &List{nil, nil, 0}
}
