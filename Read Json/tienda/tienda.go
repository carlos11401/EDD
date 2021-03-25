package tienda

type Store struct {
	Id   int
	Name string
	Age  int
}

func NewStore(id int, name string, age int) *Store {
	return &Store{id, name, age}
}
