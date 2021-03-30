package Grafo

import (
	"../Estructura"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

func GenerarMatriz(matriz *Estructura.Matriz) {
	accumInfo := "digraph G{ \n" +
		"node[shape=box];\n" +
		" Mt[ label = \"Matrix\", style = filled, fillcolor = firebrick1, group = 1 ];\n\n"

	idRowHead := ""
	linkRowHead := ""

	idColHead := ""
	linkColHead := ""

	idNodes := ""
	linksNodes := ""

	alignmentCol := " {rank=same;"
	//alignmentRow := " {rank=same;"
	countGroup := 1
	auxCountGroup := 2
	//************* ROAM ROW HEAD *********************
	nodeRowHead := matriz.EncFila.Primero
	for nodeRowHead != nil {
		if nodeRowHead.Siguiente != nil {
			// -------------- for ROW HEAD ----------------
			// id, style, color and group of Rows
			idRowHead += " \"" + fmt.Sprint(&nodeRowHead.Id) + "\"[label = \"" + strconv.Itoa(nodeRowHead.Id) + "\"" +
				" style = filled, fillcolor = bisque1, group = " + strconv.Itoa(countGroup) + " ];\n"
			// double links rows between nodeRows
			linkRowHead += " \"" + fmt.Sprint(&nodeRowHead.Id) + "\" -> \"" + fmt.Sprint(&nodeRowHead.Siguiente.Id) + "\";\n" +
				" \"" + fmt.Sprint(&nodeRowHead.Siguiente.Id) + "\" -> \"" + fmt.Sprint(&nodeRowHead.Id) + "\";\n"
		} else
		// -------------- for last ROW HEAD ----------------
		{
			idRowHead += " \"" + fmt.Sprint(&nodeRowHead.Id) + "\"[label = \"" + strconv.Itoa(nodeRowHead.Id) + "\"" +
				" style = filled, fillcolor = bisque1, group = " + strconv.Itoa(countGroup) + " ];\n"
		}
		//------------- for nodes of Rows --------------
		nodeRow := nodeRowHead.Acceso
		// for alignment of ROW
		alignmentRow := " { rank=same;\"" + fmt.Sprint(&nodeRowHead.Id) + "\";"
		for nodeRow != nil {
			// ---------- labels and group on ROWS ------------
			idNodes += " \"" + fmt.Sprint(&nodeRow.Valor) + "\"[label = \"" + nodeRow.Valor + "\""
			// for alignment of ROW
			alignmentRow += "\"" + fmt.Sprint(&nodeRow.Valor) + "\";"
			// to search in what GROUP is nodeRow
			auxColHead := matriz.EncColumna.Primero
			auxCountGroup = 2
			for auxColHead != nil {
				if auxColHead.Id == nodeRow.Columna {
					idNodes += " group = " + strconv.Itoa(auxCountGroup) + " ];\n"
					break
				}
				auxCountGroup++
				auxColHead = auxColHead.Siguiente
			}
			// ---------- links of node -----------

			if nodeRow.Izquierda != nil {
				linksNodes += " \"" + fmt.Sprint(&nodeRow.Valor) + "\" -> \"" + fmt.Sprint(&nodeRow.Izquierda.Valor) + "\";\n"
			} else {
				linksNodes += " \"" + fmt.Sprint(&nodeRowHead.Id) + "\" -> \"" + fmt.Sprint(&nodeRow.Valor) + "\";\n"
			}
			if nodeRow.Arriba != nil {
				linksNodes += " \"" + fmt.Sprint(&nodeRow.Valor) + "\" -> \"" + fmt.Sprint(&nodeRow.Arriba.Valor) + "\";\n"
			} else {
				auxCol := matriz.EncColumna.Primero
				for auxCol != nil {
					if auxCol.Id == nodeRow.Columna {
						linksNodes += " \"" + fmt.Sprint(&auxCol.Id) + "\" -> \"" + fmt.Sprint(&nodeRow.Valor) + "\";\n"
						break
					}
					auxCol = auxCol.Siguiente
				}
			}
			if nodeRow.Abajo != nil {
				if nodeRow.Abajo.Abajo != nil {
					linksNodes += " \"" + fmt.Sprint(&nodeRow.Valor) + "\" -> \"" + fmt.Sprint(&nodeRow.Abajo.Valor) + "\";\n"
				}
			}
			if nodeRow.Derecha != nil {
				linksNodes += " \"" + fmt.Sprint(&nodeRow.Valor) + "\" -> \"" + fmt.Sprint(&nodeRow.Derecha.Valor) + "\";\n"
			}
			nodeRow = nodeRow.Derecha
		}
		linksNodes += "\n" + alignmentRow + " };\n"
		nodeRowHead = nodeRowHead.Siguiente
	}
	//
	resolveProblem := ""
	linksProblems := ""
	//
	auxCol := matriz.EncColumna.Ultimo.Acceso
	for auxCol != nil {
		if auxCol.Abajo == nil && auxCol.Arriba != nil {
			linksProblems += "\"" + fmt.Sprint(&auxCol.Arriba.Valor) + "\""
			last, _ := Estructura.GetIndexAndRow(matriz, auxCol.Fila)
			penultimate, _ := Estructura.GetIndexAndRow(matriz, auxCol.Arriba.Fila)
			rest := last - penultimate
			if rest > 1 {
				for i := penultimate; i < last-1; i++ {
					// create auxiliary node
					accumInfo += " e" + strconv.Itoa(i) + "[ shape = point, width = 0 ];\n"
					//
					auxRow := Estructura.GetRowForIndex(matriz, i+1)
					resolveProblem += "\n {rank=same;\"" + fmt.Sprint(&auxRow.Id) + "\";\"e" + strconv.Itoa(i) + "\";"
					linksProblems += "->\"e" + strconv.Itoa(i) + "\""
				}
				linksProblems += "[ dir = none ];\n" +
					"\"e" + strconv.Itoa(last-2) + "\"->\"" + fmt.Sprint(&auxCol.Valor) + "\""
				resolveProblem += "}\n" + linksProblems
				accumInfo += "\n"
			}
			break
		} else if auxCol.Abajo == nil && auxCol.Arriba == nil {

		}
		auxCol = auxCol.Abajo
	}
	/*if matriz.EncFila.Size>2{
		resolveProblem = " {rank=same;\""+fmt.Sprint(&matriz.EncFila.Ultimo.Anterior.Anterior.Id)+"\";\"e0\";}\n"+
			" {rank=same;\""+fmt.Sprint(&matriz.EncFila.Ultimo.Anterior.Id)+"\";\"e1\";}\n"
		auxCol := matriz.EncColumna.Ultimo.Acceso
		for auxCol != nil{
			if auxCol.Abajo == nil{
				resolveProblem += " \""+fmt.Sprint(&auxCol.Arriba.Valor)+"\" -> \"e0\" -> \"e1\"[dir=none];\n"
				break
			}
			auxCol = auxCol.Abajo
		}

		auxRow := matriz.EncFila.Ultimo.Acceso
		for auxRow != nil{
			if auxRow.Derecha == nil{
				resolveProblem += " \"e1\" -> \""+fmt.Sprint(&auxRow.Valor)+"\";"
				break
			}
			auxRow = auxRow.Derecha
		}
	}*/
	countGroup++
	//****************** ROAM COL HEAD ******************
	nodeColHead := matriz.EncColumna.Primero
	if nodeColHead != nil {
		alignmentCol += "\"Mt\";"
		for nodeColHead.Siguiente != nil {
			idColHead += " \"" + fmt.Sprint(&nodeColHead.Id) + "\"[label=\"" + strconv.Itoa(nodeColHead.Id) + "\"" +
				" style = filled, fillcolor = bisque1, group = " + strconv.Itoa(countGroup) + " ];\n"
			linkColHead += " \"" + fmt.Sprint(&nodeColHead.Id) + "\" -> \"" + fmt.Sprint(&nodeColHead.Siguiente.Id) + "\";\n" +
				" \"" + fmt.Sprint(&nodeColHead.Siguiente.Id) + "\" -> \"" + fmt.Sprint(&nodeColHead.Id) + "\";\n"
			alignmentCol += "\"" + fmt.Sprint(&nodeColHead.Id) + "\";"
			nodeColHead = nodeColHead.Siguiente
			countGroup++
		}
		alignmentCol += " \"" + fmt.Sprint(&nodeColHead.Id) + "\";};\n\n"
		idColHead += " \"" + fmt.Sprint(&nodeColHead.Id) + "\"[label=\"" + strconv.Itoa(nodeColHead.Id) + "\"" +
			"  style = filled, fillcolor = bisque1, group = " + strconv.Itoa(countGroup) + " ];\n\n"
	}
	linkColHead += " \"Mt\" -> \"" + fmt.Sprint(&matriz.EncColumna.Primero.Id) + "\";\n" +
		" \"Mt\" -> \"" + fmt.Sprint(&matriz.EncFila.Primero.Id) + "\";\n"

	accumInfo += alignmentCol + idColHead + idRowHead + linkColHead + linkRowHead + idNodes + linksNodes + resolveProblem + "\n\n" + "\n}\n"
	path := "grafo.dot"
	//SE ESCRIBE EL ARCHIVO .DOT
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
		fmt.Println("Se ha creado un archivo")
	}

	var file, err2 = os.OpenFile(path, os.O_RDWR, 0644)
	if existeError(err2) {
		return
	}
	defer file.Close()

	//SE ESCRIBE EN ARCHIVO
	_, err = file.WriteString(accumInfo)
	if existeError(err) {
		return
	}

	// Salva los cambios
	err = file.Sync()
	if existeError(err) {
		return
	}

	fmt.Println("Archivo actualizado existosamente.")

	//PARTE EN DONDE GENERO EL GRAFO
	path2, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path2, "-Tpng", "grafo.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("grafo.png", cmd, os.FileMode(mode))
}

func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
