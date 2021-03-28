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

	//Se crea un newNode nodo con el valor ingresado
	newNode := New_Nodo(fila, columna, valor)
	boolNewRow := false
	boolNewCol := false
	//primero se busca el encabezado fila
	encFila := BuscarEncabezado(fila, matriz.EncFila)
	encColumna := BuscarEncabezado(columna, matriz.EncColumna)
	if encFila == nil { //significa que no se encontro
		//se crea el encabezado con su enlace al newNode Nodo
		encFila = New_NodoEzdo(fila, newNode)
		InsertarEncabezado(encFila, matriz.EncFila) //se insertar en el encabezado de filas
		boolNewRow = true
	}
	if encColumna == nil {
		encColumna = New_NodoEzdo(columna, newNode)
		InsertarEncabezado(encColumna, matriz.EncColumna)
		boolNewCol = true
	}
	// ******************** Row && newCol ********************
	// if there's a new Col and Row is one of the matrix
	if boolNewRow == false && boolNewCol == true {
		if encColumna == matriz.EncColumna.Primero {
			newNod := encColumna.Acceso
			encFila.Acceso.Izquierda = newNod
			newNod.Derecha = encFila.Acceso
			encFila.Acceso = newNod
		} else if encColumna == matriz.EncColumna.Ultimo {
			newNod := encColumna.Acceso
			aux := encFila.Acceso
			for aux != nil {
				if aux.Derecha == nil {
					aux.Derecha = newNod
					newNod.Izquierda = aux
					break
				}
				aux = aux.Derecha
			}
		} else {
			newNod := encColumna.Acceso
			aux := encFila.Acceso
			for aux != nil {
				if aux.Columna < newNod.Columna && newNod.Columna < aux.Derecha.Columna {
					// link with right node
					aux.Derecha.Izquierda = newNod
					newNod.Derecha = aux.Derecha
					// link with left node
					aux.Derecha = newNod
					newNod.Izquierda = aux
					break
				}
				aux = aux.Derecha
			}
		}
	} else
	// ******************** newRow && Col ********************
	// if there's a new Row and Col is one of the matrix
	if boolNewRow == true && boolNewCol == false {
		if encFila == matriz.EncFila.Primero {
			newNod := encFila.Acceso
			// link new node with down node
			encColumna.Acceso.Arriba = newNod
			newNod.Abajo = encColumna.Acceso
			// access is equal at the new node
			encColumna.Acceso = newNod
		} else if encFila == matriz.EncFila.Ultimo {
			newNod := encFila.Acceso
			aux := encColumna.Acceso
			for aux != nil {
				if aux.Abajo == nil {
					aux.Abajo = newNod
					newNod.Arriba = aux
					break
				}
				aux = aux.Abajo
			}
		} else {
			newNode := encFila.Acceso
			// roam col
			aux := encColumna.Acceso
			for aux != nil {
				if aux.Fila < newNode.Fila && newNode.Fila < aux.Abajo.Fila {
					// link with down node
					aux.Abajo.Arriba = newNode
					newNode.Abajo = aux.Abajo
					// link with left node
					aux.Abajo = newNode
					newNode.Arriba = aux
					break
				}
				aux = aux.Abajo
			}
		}
	} else
	// ******************** Row && Col ********************
	if boolNewCol == false && boolNewRow == false {
		// --------- Move Pointers of Rows -----------
		if newNode.Columna < encFila.Acceso.Columna {
			// insert newNode how first in the list of encFila
			newNode.Derecha = encFila.Acceso
			encFila.Acceso.Izquierda = newNode
			encFila.Acceso = newNode
		} else {
			actualNodeRow := encFila.Acceso
			for actualNodeRow != nil {
				// if in is because we have to insert in the last of list encFila
				if actualNodeRow.Derecha == nil {
					actualNodeRow.Derecha = newNode
					newNode.Izquierda = actualNodeRow
					break
				} else
				// find position to insert in middle of list of encFila
				if actualNodeRow.Columna < newNode.Columna && newNode.Columna < actualNodeRow.Derecha.Columna {
					// link newNode with node at the right
					newNode.Derecha = actualNodeRow.Derecha
					actualNodeRow.Derecha.Izquierda = newNode
					// link newNode with node at the left
					newNode.Izquierda = actualNodeRow
					actualNodeRow.Derecha = newNode
					break
				}
				actualNodeRow = actualNodeRow.Derecha
			}
		}
		// --------- Move Pointers of Cols -----------
		if newNode.Fila < encColumna.Acceso.Fila {
			// insert newNode how first in the list of encColumna
			newNode.Abajo = encColumna.Acceso
			encColumna.Acceso.Arriba = newNode
			encColumna.Acceso = newNode
		} else {
			actualNodeCol := encColumna.Acceso
			for actualNodeCol != nil {
				// if in here is because we have to insert in the last of list encColumna
				if actualNodeCol.Abajo == nil {
					actualNodeCol.Abajo = newNode
					newNode.Arriba = actualNodeCol
					break
				} else
				// find position to insert in middle of list of encColumna
				if actualNodeCol.Fila < newNode.Fila && newNode.Fila < actualNodeCol.Abajo.Fila {
					// link newNode with node that is down
					newNode.Abajo = actualNodeCol.Abajo
					actualNodeCol.Abajo.Arriba = newNode
					// link newNode with node that is above
					newNode.Arriba = actualNodeCol
					actualNodeCol.Abajo = newNode
					break
				}
				actualNodeCol = actualNodeCol.Abajo
			}
		}
	}
}
