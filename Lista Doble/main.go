package main

import (
	"./Lista"
	"fmt"
)

func main() {
	fmt.Println("----------- LISTA ----------")
	doubleList := Lista.NewList()
	a := Lista.Node{Nombre: "Carlos", Edad: 24}
	b := Lista.Node{Nombre: "Javier", Edad: 25}
	c := Lista.Node{Nombre: "Jose", Edad: 26}
	doubleList.Insert(&a)
	doubleList.Insert(&b)
	doubleList.Insert(&c)

	doubleList.Print()
}
