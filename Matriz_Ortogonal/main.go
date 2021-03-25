package main

import (
	"./Estructura"
	"fmt"
	"strconv"
)

// ptr := fmt.Sprint(&matriz) //para obtener la direccion de un nodo

func main() {
	fmt.Println("hola mundo ")
	matriz := Estructura.New_Matriz()
	/*Estructura.Insertar_Matriz(matriz,-2,3,"20")
	Estructura.Insertar_Matriz(matriz,-2,2,"10")
	Estructura.Insertar_Matriz(matriz,-2,5,"30")
	Estructura.Insertar_Matriz(matriz,-2,6,"40")
	Estructura.Insertar_Matriz(matriz,-2,4,"25")*/

	Estructura.Insertar_Matriz(matriz, 2, 2, "30")
	Estructura.Insertar_Matriz(matriz, 1, 2, "20")
	Estructura.Insertar_Matriz(matriz, 0, 2, "50")
	Estructura.Insertar_Matriz(matriz, 3, 2, "40")
	Estructura.Insertar_Matriz(matriz, 4, 2, "70")
	//Estructura.Insertar_Matriz(matriz,-5,2,"10")
	//Estructura.Insertar_Matriz(matriz,-1,2,"50")
	nodeRow := matriz.EncFila.Primero
	fmt.Println("Filas")
	strRow := ""
	for nodeRow != nil {
		aux := nodeRow.Acceso
		strRow += "Row: " + strconv.Itoa(aux.Fila) + " --> "
		for aux != nil {
			strRow += aux.Valor + " --> "
			aux = aux.Derecha
		}
		strRow += "\n"
		nodeRow = nodeRow.Siguiente
	}
	fmt.Println(strRow)
	fmt.Println("Col")
	strCol := ""
	nodeCol := matriz.EncColumna.Primero
	for nodeCol != nil {
		aux := nodeCol.Acceso
		strCol += "Col: " + strconv.Itoa(aux.Columna) + " --> "
		for aux != nil {
			strCol += aux.Valor + " --> "
			aux = aux.Abajo
		}
		strCol += "\n"
		nodeCol = nodeCol.Siguiente
	}
	fmt.Println(strCol)
	//Grafo.GenerarMatriz(matriz)
}
