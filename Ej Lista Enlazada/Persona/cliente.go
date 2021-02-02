package Persona

type Cliente struct {
	Id   int
	Name string
	Age  int
}

func New_Client(id int, name string, age int) *Cliente {
	return &Cliente{id, name, age}
}
