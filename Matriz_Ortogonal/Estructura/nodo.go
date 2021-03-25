package Estructura

type Nodo struct {
	Valor string
	Fila int
	Columna int
	Izquierda *Nodo
	Derecha *Nodo
	Arriba *Nodo
	Abajo *Nodo
}


func New_Nodo(fila int , columna int, valor string) *Nodo{
	return &Nodo{valor,fila,columna,nil,nil,nil,nil}
}