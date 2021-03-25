package Estructura

import (
	// "fmt"
)
type Matriz struct {
	EncFila *Encabezado 
	EncColumna *Encabezado
}

func New_Matriz() *Matriz{
	return &Matriz{New_Encabezado(),New_Encabezado()}
}

func Insertar_Matriz(matriz *Matriz,fila int, columna int, valor string) {

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
			newNode := encColumna.Acceso
			aux := encFila.Acceso
			for aux != nil {
				if aux.Derecha == nil {
					aux.Derecha = newNode
					newNode.Izquierda = aux
					break
				}
				aux = aux.Derecha
			}
		} else {
			newNode := encColumna.Acceso
			aux := encFila.Acceso
			for aux != nil {
				if aux.Columna < newNode.Columna && newNode.Columna < aux.Derecha.Columna {
					// link with right node
					aux.Derecha.Izquierda = newNode
					newNode.Derecha = aux.Derecha
					// link with left node
					aux.Derecha = newNode
					newNode.Izquierda = aux
					break
				}
				aux = aux.Derecha
			}
		}
	}else
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
			newNode := encFila.Acceso
			aux := encColumna.Acceso
			for aux != nil {
				if aux.Abajo == nil {
					aux.Abajo = newNode
					newNode.Arriba = aux
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
	}else
	// ******************** Row && Col ********************
	if boolNewCol == false && boolNewRow == false{
		// ------- COL last and ROW last -------
		if encColumna == matriz.EncColumna.Ultimo && encFila == matriz.EncFila.Ultimo{
			// Roaming Row
			actualNodeRow := encFila.Acceso
			for actualNodeRow != nil{
				if actualNodeRow.Derecha == nil{
					// linking last node of ROW with newNodo
					actualNodeRow.Derecha = newNode
					newNode.Izquierda = actualNodeRow
					break
				}
				actualNodeRow = actualNodeRow.Derecha
			}
			// Roaming Col
			actualNodeCol := encColumna.Acceso
			for actualNodeCol != nil{
				if actualNodeCol.Abajo == nil{
					// linking last node of Col with newNode
					actualNodeCol.Abajo = newNode
					newNode.Arriba = actualNodeRow
					break
				}
				actualNodeCol = actualNodeCol.Abajo
			}
		}else
		// ------- COL first and ROW first -------
		if encColumna == matriz.EncColumna.Primero && encFila == matriz.EncFila.Primero{
			// linking newNode with Node Access of ROW
			encFila.Acceso.Izquierda = newNode
			newNode.Derecha = encFila.Acceso
			// linking newNode with Node Access of COL
			encColumna.Acceso.Arriba = newNode
			newNode.Abajo = encColumna.Acceso
			// --- How now newNode will be the First ---
			// access of Row is newNode
			encFila.Acceso = newNode
			// access of Col is newNode
			encColumna.Acceso = newNode
		}else
		// ------- COL first and ROW first -------
		if encColumna == matriz.EncColumna.Primero{
			//
			actualNodeCol := encColumna.Acceso
			for actualNodeCol != nil{
				if actualNodeCol.Fila < newNode.Fila && newNode.Fila < actualNodeCol.Abajo.Fila{

				}
				actualNodeCol = actualNodeCol.Abajo
			}
		}
	}
}