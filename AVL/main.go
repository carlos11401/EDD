package main

func main() {
	a := NewTree()
	for i := 0; i < 20; i++ {
		a.Insert(i)
	}
	GenerateGraph(a)
}
