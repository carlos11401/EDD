package Persona

type Client struct {
	Id   int
	Name string
	Age  int
}

func NewClient(id int, name string, age int) *Client {
	return &Client{id, name, age}
}
