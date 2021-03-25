package main

import "fmt"

type nodo struct {
	num int
	Siguiente, Anterior *nodo
}

type lista struct {
	cabeza *nodo
	cola   *nodo
}
func (this *lista) Insertar(nuevo *nodo) {
	if this.cabeza == nil {
		this.cabeza = nuevo
		this.cola = nuevo
	} else {
		this.cola.Siguiente = nuevo
		nuevo.Anterior = this.cola
		this.cola = nuevo
	}
}
func (this *lista)Order(){
	var siguiente *nodo
	var temp int
	actual := this.cabeza
	for actual.Siguiente != nil{
		siguiente =  actual.Siguiente
		for siguiente != nil{
			if actual.num > siguiente.num{
				temp = siguiente.num
				siguiente.num = actual.num
				actual.num = temp
			}
			siguiente = siguiente.Siguiente
		}
		actual = actual.Siguiente
		siguiente = actual.Siguiente
	}
}
func main() {
	list := lista{nil, nil}
	num1 := nodo{21,nil,nil}
	num2 := nodo{43,nil,nil}
	num3 := nodo{1,nil,nil}
	num4 := nodo{5,nil,nil}
	num5 := nodo{6,nil,nil}
	num6 := nodo{76,nil,nil}
	list.Insertar(&num1)
	list.Insertar(&num2)
	list.Insertar(&num3)
	list.Insertar(&num4)
	list.Insertar(&num5)
	list.Insertar(&num6)
	list.Order()
	aux := list.cabeza
	for aux != nil{
		fmt.Print(aux.num," -> ")
		aux = aux.Siguiente
	}
}
