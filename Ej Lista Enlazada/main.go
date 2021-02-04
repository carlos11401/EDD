package main

import (
	"./Estructura"
	"./Persona"
	"fmt"
)

func main() {
	fmt.Println("///////////////// WELCOME ////////////////")
	// create list
	var list *Estructura.List = Estructura.NewList()

	var selected string
	for selected != "close" {
		fmt.Print("--> ")
		fmt.Scan(&selected)

		switch selected {
		case "insert":
			fmt.Println("	>>>>>>>>>> INSERT <<<<<<<<<<")
			var name string
			var id, age int
			fmt.Print("--> name: ")
			fmt.Scanln(&name)
			fmt.Print("--> id: ")
			fmt.Scanln(&id)
			fmt.Print("--> age: ")
			fmt.Scanln(&age)

			var cli *Persona.Client = Persona.NewClient(id, name, age)
			Estructura.Insert(cli, list)
			Estructura.Print(list)

		case "delete":
			Estructura.Delete(list)
			Estructura.Print(list)
		case "print":
			Estructura.Print(list)
		}
	}
}
