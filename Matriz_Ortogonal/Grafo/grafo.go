package Grafo

import (
	"../Estructura"
	"fmt"
	"strconv"
	"os"
	"io/ioutil"
	"os/exec"
)

func GenerarMatriz(matriz * Estructura.Matriz){
	acumInfo := "digraph G{ \n"+
	"node[shape=box, style=filled, color=deepskyblue3];\n"+
	"edge[color=black];\n"+
	"rankdir=UD;\n"

	idCabeceraFila :=""
	cabeceraFila :=""
	
	idCabeceraCol :=""
	cabeceraCol :=""

	alineacionCol := "{rank=min; Matriz;"

	//************* RECORRIDO DE FILAS *********************
	eFila := matriz.EncFila.Primero;

	if eFila != nil{
		cabeceraFila += "Matriz ->" + "\""+fmt.Sprint(&eFila.Id) + "\";\n"
		for eFila.Siguiente != nil{
			idCabeceraFila += "\"" +fmt.Sprint(&eFila.Id) + "\"[label = \"" + strconv.Itoa(eFila.Id) + "\"];\n"
			cabeceraFila += "\"" + fmt.Sprint(&eFila.Id) + "\" -> \"" +  fmt.Sprint(&eFila.Siguiente.Id) + "\";\n"
			eFila = eFila.Siguiente
		}
		idCabeceraFila += "\"" + fmt.Sprint(&eFila.Id) + "\"[label = \"" + strconv.Itoa(eFila.Id) + "\"];\n"
	
	}

	//****************** RECORRIDO DE COLUMNAS ******************
	eCol := matriz.EncColumna.Primero

	if eCol != nil{
		cabeceraCol += "Matriz ->" + "\"" + fmt.Sprint(&eCol.Id) + "\";\n"
		
		for eCol.Siguiente != nil{
			alineacionCol += "\"" + fmt.Sprint(&eCol.Id) + "\";"
			idCabeceraCol += "\"" + fmt.Sprint(&eCol.Id) + "\"[label=\"" + strconv.Itoa(eCol.Id) + "\"];\n"
			cabeceraCol += "\"" + fmt.Sprint(&eCol.Id) + "\" -> \"" + fmt.Sprint(&eCol.Siguiente.Id) + "\";\n"
			eCol = eCol.Siguiente
		}

		alineacionCol += "\"" + fmt.Sprint(&eCol.Id) + "\";};\n\n"
		idCabeceraCol += "\"" + fmt.Sprint(&eCol.Id) + "\"[label=\"" +  strconv.Itoa(eCol.Id) +"\"];\n"
		
		//Graficar los punteros hacia atras 
	}

	acumInfo += alineacionCol + idCabeceraCol +  idCabeceraFila + cabeceraCol + cabeceraFila + "\n\n" + "\n}\n"

	path := "grafo.dot"
	//SE ESCRIBE EL ARCHIVO .DOT
	var _, err = os.Stat(path)
	if os.IsNotExist(err){
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
	 _, err = file.WriteString(acumInfo)
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
	cmd, _ := exec.Command(path2, "-Tpng","grafo.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("grafo.png",cmd, os.FileMode(mode))
}

func existeError(err error) bool{
	if err != nil{
		fmt.Println(err.Error())
	}
	return (err != nil)
}