package main

import (
	"./Estructura"
	"./Persona"
)

func main() {
	var cli *Persona.Cliente = Persona.New_Client(3, "cli 1", 20)
	var list *Estructura.List = Estructura.New_List()
	Estructura.Insert(cli, list)
	Estructura.Print(list)
}
