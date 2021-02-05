package main

import (
	"./Lista"
	"fmt"
)

func main() {
	fmt.Println("----------- LISTA ----------")
	doubleList := Lista.NewList()
	a := Lista.Node{"Carlos", 24, nil, nil}
	b := Lista.Node{"Javier", 25, nil, nil}
	c := Lista.Node{"Jose", 26, nil, nil}
	doubleList.Insert(&a)
	doubleList.Insert(&b)
	doubleList.Insert(&c)
}
