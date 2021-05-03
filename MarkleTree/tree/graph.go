package tree

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func (thisTree *Tree) GenerateGraph() {
	dotStructure := "digraph G{\nnode [shape=rectangle];\n"
	accum := ""

	if thisTree.root != nil {
		RoamTree(&thisTree.root, &accum)
	}

	dotStructure += accum + "\n}\n"

	path := "grafo.dot"
	//SE ESCRIBE EL ARCHIVO .DOT
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if existError(err) {
			return
		}
		defer file.Close()
		fmt.Println("Se ha creado un archivo")
	}

	var file, err2 = os.OpenFile(path, os.O_RDWR, 0644)
	if existError(err2) {
		return
	}
	defer file.Close()

	//SE ESCRIBE EN ARCHIVO
	_, err = file.WriteString(dotStructure)
	if existError(err) {
		return
	}

	// Salva los cambios
	err = file.Sync()
	if existError(err) {
		return
	}

	fmt.Println("Archivo actualizado existosamente.")

	//PARTE EN DONDE GENERO EL GRAFO
	path2, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path2, "-Tpng", "grafo.dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("grafo.png", cmd, os.FileMode(mode))

}
func RoamTree(actual **Node, acum *string) {
	if *actual != nil {
		//SE OBTIENE INFORMACION DEL NODO ACTUAL
		text, countRows := "", 0
		// just get value of 20 chars for row
		for countRows < 4 && len((*actual).value) > (countRows+1)*20 {
			text += (*actual).value[countRows*20:(countRows+1)*20] + "\n"
			countRows++
		}
		if countRows == 0 {
			text = (*actual).value
		}
		*acum += "\"" + fmt.Sprint(&(*actual)) + "\"[label=\"" + text + "\"];\n"
		//VIAJAMOS A LA SUBRAMA IZQ
		if (*actual).left != nil {
			*acum += "\"" + fmt.Sprint(&(*actual)) + "\" -> \"" + fmt.Sprint(&(*actual).left) + "\";\n"
		}
		//VIAJAMOS A LA SUBRAMA DER
		if (*actual).right != nil {
			*acum += "\"" + fmt.Sprint(&(*actual)) + "\" -> \"" + fmt.Sprint(&(*actual).right) + "\";\n"
		}
		RoamTree(&(*actual).left, acum)
		RoamTree(&(*actual).right, acum)

	}
}

func existError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return err != nil
}
