package main

import "fmt"

func main() {
	x := 25
	y := 100
	fmt.Print(MDC(x,y))
}
func MDC (x,y int)int{
	var result int
	if x == y{
		result = x
	}else if x>y {
		result = MDC(x-y,y)
	}else if y>x {
		result = MDC(x,y-x)
	}
	return result
}