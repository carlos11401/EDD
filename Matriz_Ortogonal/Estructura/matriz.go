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

	//Se crea un nuevo nodo con el valor ingresado
	nuevo := New_Nodo(fila, columna, valor)
	boolNewRow := false
	boolNewCol := false
	//********************** INSERCION EN FILAS ***********************

	//primero se busca el encabezado fila
	encFila := BuscarEncabezado(fila, matriz.EncFila)
	encColumna := BuscarEncabezado(columna, matriz.EncColumna)
	if encFila == nil { //significa que no se encontro
		//se crea el encabezado con su enlace al nuevo Nodo
		encFila = New_NodoEzdo(fila, nuevo)
		InsertarEncabezado(encFila, matriz.EncFila) //se insertar en el encabezado de filas
		boolNewRow = true
	}
	if encColumna == nil{
		encColumna = New_NodoEzdo(columna,nuevo)
		InsertarEncabezado(encColumna, matriz.EncColumna)
		boolNewCol = true
	}
	// if there's a new Col and Row is one of the matrix
	if boolNewRow==false && boolNewCol==true{
		if encColumna == matriz.EncColumna.Primero{
			newNode := encColumna.Acceso
			encFila.Acceso.Izquierda = newNode
			newNode.Derecha = encFila.Acceso
			encFila.Acceso = newNode
		}else if encColumna == matriz.EncColumna.Ultimo{
			newNode := encColumna.Acceso
			aux := encFila.Acceso
			for aux != nil{
				if aux.Derecha == nil{
					aux.Derecha = newNode
					newNode.Izquierda = aux
					break
				}
				aux = aux.Derecha
			}
		}else{
			newNode := encColumna.Acceso
			aux := encFila.Acceso
			for aux != nil{
				if aux.Columna<newNode.Columna && newNode.Columna<aux.Derecha.Columna{
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
	}
	// if there's a new Row and Col is one of the matrix
	if boolNewRow==true && boolNewCol==false{
		if encFila == matriz.EncFila.Primero{
			newNode := encFila.Acceso
			// link new node with down node
			encColumna.Acceso.Arriba = newNode
			newNode.Abajo = encColumna.Acceso
			// access is equal at the new node
			encColumna.Acceso = newNode
		}else if encFila == matriz.EncFila.Ultimo{
			newNode := encFila.Acceso
			aux := encColumna.Acceso
			for aux != nil{
				if aux.Abajo == nil{
					aux.Abajo = newNode
					newNode.Arriba = aux
					break
				}
				aux = aux.Abajo
			}
		}else{
			newNode := encFila.Acceso
			// roam col
			aux := encColumna.Acceso
			for aux != nil{
				if aux.Fila<newNode.Fila && newNode.Fila<aux.Abajo.Fila{
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
	}
}
	// ******************** INSERCION EN COLUMNAS **************************