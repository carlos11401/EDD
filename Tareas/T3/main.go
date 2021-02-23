package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type nodo struct {
	nombre, apellido, apodo, favoritos string
	Siguiente, Anterior                *nodo
}

type lista struct {
	cabeza *nodo
	cola   *nodo
}

func createGraph(list lista) {
	acum := "digraph G{\n rankdir = LR;\nnode [shape=box]; \n compound=true; \n"
	content := ""
	link := ""
	if list.cabeza != nil {
		node := list.cabeza
		nodeGraph := ""
		countNode := 0

		for node.Siguiente != nil {
			nodeGraph += strconv.Itoa(countNode) + "[label=\"Apellido : " + node.apellido + "\n"
			nodeGraph += "Nombre : " + node.nombre + "\n"
			nodeGraph += "Apodo : " + node.apodo + "\n"
			nodeGraph += "Favoritos : " + node.favoritos + "\"];\n"

			link += strconv.Itoa(countNode) + "->" + strconv.Itoa(countNode+1) + ";\n"
			link += strconv.Itoa(countNode+1) + "->" + strconv.Itoa(countNode) + ";\n"
			node = node.Siguiente
			countNode++
		}
		nodeGraph += strconv.Itoa(countNode) + "[label=\"Apellido : " + node.apellido + "\n"
		nodeGraph += "Nombre : " + node.nombre + "\n"
		nodeGraph += "Apodo : " + node.apodo + "\n"
		nodeGraph += "Favoritos : " + node.favoritos + "\"];\n"
		content = acum + nodeGraph + link + "\n}\n"
		createDotImage(content, 1)
	}
}
func createDotImage(content string, i int) {
	//creacion del archivo
	path := "grafo" + strconv.Itoa(i) + ".dot"
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}

	//AHORA ESCRIBIEMOS EN EL ARCHIVO

	var file, err2 = os.OpenFile(path, os.O_RDWR, 0644)
	if existeError(err2) {
		return
	}
	defer file.Close()

	//SE ESCRIBE EN ARCHIVO
	_, err = file.WriteString(content)
	if existeError(err) {
		return
	}

	// Salva los cambios
	err = file.Sync()
	if existeError(err) {
		return
	}

	//PARTE EN DONDE GENERO EL GRAFO
	path2, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path2, "-Tpng", path).Output()
	mode := int(0777)
	ioutil.WriteFile("grafo"+strconv.Itoa(i)+".png", cmd, os.FileMode(mode))
}
func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return err != nil
}
func (this *lista) Insertar(nuevo *nodo) {
	if this.cabeza == nil {
		this.cabeza = nuevo
		this.cola = nuevo
	} else {
		this.cola.Siguiente = nuevo
		nuevo.Anterior = this.cola
		this.cola = nuevo
	}
}

func main() {
	li := lista{nil, nil}
	a := nodo{"Marvin", "Martinez", "Marvin25ronal", "Jugar apex", nil, nil}
	b := nodo{"Yaiza", "Pineda", "Bambi", "Patinar", nil, nil}
	c := nodo{"Jonathan", "Lopez", "Pancho", "Comer", nil, nil}
	d := nodo{"usuario1", "bla", "bla", "Jugar apex", nil, nil}
	e := nodo{"usuario2", "bla", "bla", "Jugar apex", nil, nil}
	f := nodo{"usuario3", "sale edd", "vamos con todo", "100 en la fase 1", nil, nil}
	li.Insertar(&a)
	li.Insertar(&b)
	li.Insertar(&c)
	li.Insertar(&d)
	li.Insertar(&e)
	li.Insertar(&f)
	createGraph(li)
}
