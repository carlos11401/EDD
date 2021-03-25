package main

import (
	"./Estructura"
	"fmt"
	"strconv"
)
// ptr := fmt.Sprint(&matriz) //para obtener la direccion de un nodo 

func main(){
	matriz := Estructura.New_Matriz()
	Estructura.Insertar_Matriz(matriz,5,2,"1")
	Estructura.Insertar_Matriz(matriz,3,2,"2")
	Estructura.Insertar_Matriz(matriz,6,2,"3")
	Estructura.Insertar_Matriz(matriz,4,2,"4")
	nodeRow:= matriz.EncFila.Primero
	str := ""
	for nodeRow != nil{
		node := nodeRow.Acceso
		str += "Row: "+strconv.Itoa(nodeRow.Id)+" -> "
		for node != nil{
			str += node.Valor+" <--> "
			node = node.Derecha
		}
		str += "\n"
		nodeRow = nodeRow.Siguiente
	}
	nodeCol:= matriz.EncColumna.Primero
	for nodeCol != nil{
		node := nodeCol.Acceso
		str += "Col: "+strconv.Itoa(nodeCol.Id)+" -> "
		for node != nil{
			str += node.Valor+" <--> "
			node = node.Abajo
		}
		str += "\n"
		nodeCol = nodeCol.Siguiente
	}
	fmt.Println(str)
	//Grafo.GenerarMatriz(matriz)
}