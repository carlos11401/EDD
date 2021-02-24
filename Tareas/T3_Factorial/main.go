package main

import "fmt"

func main() {
	var num int
	fmt.Print("Ingresa un Numero: ")
	fmt.Scanln(&num)
	fmt.Print("El factorial de ", num, " es ", factorial(num))
}

func factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}
