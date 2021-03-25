package Estructura

import (
// "fmt"
)

type Matriz struct {
	EncFila    *Encabezado
	EncColumna *Encabezado
}

func New_Matriz() *Matriz {
	return &Matriz{New_Encabezado(), New_Encabezado()}
}

func Insertar_Matriz(matriz *Matriz, fila int, columna int, valor string) {
	//Se crea un newNodo nodo con el valor ingresado
	newNodo := New_Nodo(fila, columna, valor)
	//search in heads (col,row) if exist the head asked
	encFila := BuscarEncabezado(fila, matriz.EncFila)
	encColumna := BuscarEncabezado(columna, matriz.EncColumna)
	boolFilaNew := false
	boolColNew := false
	// insert if we have to create a new head or there's yet in head Row
	if encFila == nil {
		/*significa que no se encontro
		  se crea el encabezado con su enlace al newNodo Nodo*/
		encFila = New_NodoEzdo(fila, newNodo)
		//se insertar en el encabezado de filas
		InsertarEncabezado(encFila, matriz.EncFila)
		boolFilaNew = true
	}
	// insert if we have to create a new head or there's yet in head Col
	if encColumna == nil {
		encColumna = New_NodoEzdo(columna, newNodo)
		InsertarEncabezado(encColumna, matriz.EncColumna)
		boolColNew = true
	}
	if boolFilaNew == false && boolColNew == true {
		//********************** ROW ***********************
		if encColumna == matriz.EncColumna.Primero {
			newNode := matriz.EncColumna.Primero.Acceso
			// insert new node how First (ACCESS) in EncFila
			encFila.Acceso.Izquierda = newNode
			newNode.Derecha = encFila.Acceso
			matriz.EncFila.Primero.Acceso = newNode
			// when a row repeat and col head is last
		} else if encColumna == matriz.EncColumna.Ultimo {
			// insert new node how First (ACCESS) in EncFila
			temp := encFila.Acceso
			for temp != nil { // roam access, where is the nodes
				if temp.Derecha == nil {
					temp.Derecha = matriz.EncColumna.Ultimo.Acceso
					matriz.EncColumna.Ultimo.Acceso.Izquierda = temp
					break
				}
				temp = temp.Derecha
			}
			// when a row repeat and col head don't is first neither last
		} else if encFila != nil {
			temp := encColumna.Anterior.Acceso
			// roam access of before head col of new Col
			for temp != nil {
				// finding Row where we want to insert a new Node
				if encFila.Id == temp.Fila {
					// linking with node after to new node
					temp.Derecha.Izquierda = encColumna.Acceso
					encColumna.Acceso.Derecha = temp.Derecha
					// linking with node before to new node
					temp.Derecha = encColumna.Acceso
					encColumna.Acceso.Izquierda = temp
					break
				}
				temp = temp.Abajo
			}
		}
	}
	if boolFilaNew == true && boolColNew == false {
		//********************** COL ***********************
		if encFila == matriz.EncFila.Primero {
			temp := matriz.EncFila.Primero.Acceso
			temp.Abajo = encColumna.Acceso
			encColumna.Acceso.Arriba = temp
			encColumna.Acceso = temp
		} else if encFila == matriz.EncFila.Ultimo {
			aux := encColumna.Acceso
			for aux != nil {
				if aux.Abajo == nil {
					aux.Abajo = encFila.Acceso
					encFila.Acceso.Arriba = aux
					break
				}
				aux = aux.Abajo
			}
		}
	}
}
