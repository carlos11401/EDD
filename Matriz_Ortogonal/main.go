package main

import (
	"./Estructura"
	"./Grafo"
	"fmt"
	"strconv"
)

// ptr := fmt.Sprint(&matriz) //para obtener la direccion de un nodo

func main() {
	matriz := Estructura.New_Matriz()
	Estructura.Insertar_Matriz(matriz, 1, 2, "4")
	Estructura.Insertar_Matriz(matriz, 2, 3, "5")
	Estructura.Insertar_Matriz(matriz, 1, 4, "6")
	Estructura.Insertar_Matriz(matriz, 1, 3, "9")
	//Estructura.Insertar_Matriz(matriz, 2, 5, "7")
	//Estructura.Insertar_Matriz(matriz, 1, 5, "10")
	Estructura.Insertar_Matriz(matriz, 2, 1, "11")
	Estructura.Insertar_Matriz(matriz, 1, 1, "3")
	Estructura.Insertar_Matriz(matriz, 3, 2, "20")
	Estructura.Insertar_Matriz(matriz, 3, 4, "21")
	Estructura.Insertar_Matriz(matriz, 2, 4, "22")
	Estructura.Insertar_Matriz(matriz, 4, 1, "23")
	Estructura.Insertar_Matriz(matriz, 2, 2, "33")
	//Estructura.Insertar_Matriz(matriz, 4, 5, "55")

	nodeRow := matriz.EncFila.Primero
	str := ""
	for nodeRow != nil {
		node := nodeRow.Acceso
		str += "Row:" + strconv.Itoa(nodeRow.Id) + " --::-- "
		for node != nil {
			str += node.Valor + " <--> "
			node = node.Derecha
		}
		str += "\n"
		nodeRow = nodeRow.Siguiente
	}
	str += "###############\n"
	nodeCol := matriz.EncColumna.Primero
	for nodeCol != nil {
		node := nodeCol.Acceso
		str += "Col:" + strconv.Itoa(nodeCol.Id) + " --::-- "
		for node != nil {
			str += node.Valor + " <--> "
			node = node.Abajo
		}
		str += "\n"
		nodeCol = nodeCol.Siguiente
	}
	fmt.Println(str)
	Grafo.GenerarMatriz(matriz)
}
