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
	} else
	// ******************** Row && Col ********************
	if boolNewCol == false && boolNewRow == false {
		// ------------------ CORNERS ----------------------
		// ------- COL first and ROW first -------
		if encColumna == matriz.EncColumna.Primero && encFila == matriz.EncFila.Primero {
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
		}
		// ------- COL last and ROW last -------
		if encColumna == matriz.EncColumna.Ultimo && encFila == matriz.EncFila.Ultimo {
			// Roaming Row
			actualNodeRow := encFila.Acceso
			for actualNodeRow != nil {
				if actualNodeRow.Derecha == nil {
					// linking last node of ROW with newNodo
					actualNodeRow.Derecha = newNode
					newNode.Izquierda = actualNodeRow
					break
				}
				actualNodeRow = actualNodeRow.Derecha
			}
			// Roaming Col
			actualNodeCol := encColumna.Acceso
			for actualNodeCol != nil {
				if actualNodeCol.Abajo == nil {
					// linking last node of Col with newNode
					actualNodeCol.Abajo = newNode
					newNode.Arriba = actualNodeRow
					break
				}
				actualNodeCol = actualNodeCol.Abajo
			}
		} else
		// ------- Row last and COL first -------
		if encFila == matriz.EncFila.Ultimo && encColumna == matriz.EncColumna.Primero {
			// linking access of Row with newNode
			encFila.Acceso.Izquierda = newNode
			newNode.Derecha = encFila.Acceso
			encFila.Acceso = newNode
			// Roaming col
			actualNodeCol := encColumna.Acceso
			for actualNodeCol != nil {
				if actualNodeCol.Abajo == nil {
					// linking last node of col with newNode
					actualNodeCol.Abajo = newNode
					newNode = actualNodeCol
					break
				}
				actualNodeCol = actualNodeCol.Abajo
			}
		} else
		// ------- Row first and COL last -------
		if encFila == matriz.EncFila.Primero && encColumna == matriz.EncColumna.Ultimo {
			// linking access of Col with newNode
			encColumna.Acceso.Arriba = newNode
			newNode.Abajo = encColumna.Acceso
			encColumna.Acceso = newNode
			// Roam row
			acutalNodeRow := encFila.Acceso
			for acutalNodeRow != nil {
				if acutalNodeRow.Derecha == nil {
					acutalNodeRow.Derecha = newNode
					newNode.Izquierda = acutalNodeRow
					break
				}
				acutalNodeRow = acutalNodeRow.Derecha
			}
		} else
		// ------------------  EDGES   ----------------------
		// -------  Up Edge  -------
		if encFila == matriz.EncFila.Primero {
			// Roam Row to search where insert the newNode
			actualNodeRow := encFila.Acceso
			for actualNodeRow != nil {
				if actualNodeRow.Columna < newNode.Columna && newNode.Columna < actualNodeRow.Derecha.Columna {
					// linking newNode <--> actualNode.Derecha
					actualNodeRow.Derecha.Izquierda = newNode
					newNode.Derecha = actualNodeRow.Derecha
					// linking actualNode <--> newNode
					actualNodeRow.Derecha = newNode
					newNode.Izquierda = actualNodeRow
					// linking newNode with node of down
					encColumna.Acceso.Arriba = newNode
					newNode.Abajo = encColumna.Acceso
					// change access to newNode
					encColumna.Acceso = newNode
					break
				}
				actualNodeRow = actualNodeRow.Derecha
			}
		} else
		// ------- Down Edge -------
		if encFila == matriz.EncFila.Ultimo {
			// Roam Row to search where insert the newNode
			actualNodeRow := encFila.Acceso
			for actualNodeRow != nil {
				if actualNodeRow.Derecha != nil {
					if actualNodeRow.Columna < newNode.Columna && newNode.Columna < actualNodeRow.Derecha.Columna {
						// linking newNode <--> actualNode.Derecha
						actualNodeRow.Derecha.Izquierda = newNode
						newNode.Derecha = actualNodeRow.Derecha
						// linking actualNode <--> newNode
						actualNodeRow.Derecha = newNode
						newNode.Izquierda = actualNodeRow
						// linking newNode with node of Up
						actualNodeCol := encColumna.Acceso
						for actualNodeCol != nil {
							if actualNodeCol.Abajo == nil {
								actualNodeCol.Abajo = newNode
								newNode.Arriba = actualNodeCol
								break
							}
							actualNodeCol = actualNodeCol.Abajo
						}
						break
					}
				} else if newNode.Columna < actualNodeRow.Columna {
					// linking actualNode <--> newNode
					actualNodeRow.Izquierda = newNode
					newNode.Derecha = actualNodeRow
					encFila.Acceso = newNode
					// linking newNode with node of Up
					actualNodeCol := encColumna.Acceso
					for actualNodeCol != nil {
						if actualNodeCol.Abajo == nil {
							actualNodeCol.Abajo = newNode
							newNode.Arriba = actualNodeCol
							break
						}
						actualNodeCol = actualNodeCol.Abajo
					}
					break
				} else {
					// linking actualNode <--> newNode
					actualNodeRow.Derecha = newNode
					newNode.Izquierda = actualNodeRow
					// linking newNode with node of Up
					actualNodeCol := encColumna.Acceso
					for actualNodeCol != nil {
						if actualNodeCol.Abajo == nil {
							actualNodeCol.Abajo = newNode
							newNode.Arriba = actualNodeCol
							break
						}
						actualNodeCol = actualNodeCol.Abajo
					}
					break
				}
				actualNodeRow = actualNodeRow.Derecha
			}
		} else
		// ------- Left Edge -------
		if encColumna == matriz.EncColumna.Primero {
			// Roam Col
			actualNodeCol := encColumna.Acceso
			for actualNodeCol != nil {
				if actualNodeCol.Abajo != nil {
					// search where have to insert newNode
					if actualNodeCol.Fila < newNode.Fila && newNode.Fila < actualNodeCol.Abajo.Fila {
						// linking newNode.Abajo with actualNode.Abajo
						actualNodeCol.Abajo.Arriba = newNode
						newNode.Abajo = actualNodeCol.Abajo
						// linking newNode.Arriba with actualNode
						actualNodeCol.Abajo = newNode
						newNode.Arriba = actualNodeCol
						// linking newNode.Derecha with encFila.Access
						encFila.Acceso.Izquierda = newNode
						newNode.Derecha = encFila.Acceso
						// change access to newNode
						encFila.Acceso = newNode
						break
					}
				} else {
					// linking newNode.Arriba with actualNode
					actualNodeCol.Abajo = newNode
					newNode.Arriba = actualNodeCol
					// linking newNode.Derecha with encFila.Access
					encFila.Acceso.Izquierda = newNode
					newNode.Derecha = encFila.Acceso
					// change access to newNode
					encFila.Acceso = newNode
					break
				}
				actualNodeCol = actualNodeCol.Abajo
			}
		}
	}
}
