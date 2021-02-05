package Lista

import (
	"fmt"
	"strconv"
)

type Node struct {
	Nombre string
	Edad   int
	Next   *Node
	Back   *Node
}

type List struct {
	first, last *Node
}

func NewList() *List {
	return &List{nil, nil}
}

func (this *List) Insert(new *Node) {
	if this.first == nil {
		this.first = new
		this.last = new
	} else {
		this.last.Next = new
		new.Back = this.last
		this.last = new
	}
}
func (this *Node) ToString() string {
	return "Nombre: " + this.Nombre + " Edad: " + strconv.Itoa(this.Edad)
}
func (this *List) ToString() string {
	var cadena string
	aux := this.first
	for aux != nil {
		cadena += aux.ToString()
		aux = aux.Next
	}
	return cadena
}

func (this *List) Print() {
	fmt.Print("lista--------------------")
	fmt.Println(this.ToString())
}
