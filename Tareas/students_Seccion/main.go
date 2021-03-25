package main
type List struct {
	First, Last *Node
	Count       int
}
type Node struct {
	seccion int
	listEst *ListEst
	Next, Back  *Node
}
type NodeStd struct {
	carne int
	Next, Back  *Node
}
type ListEst struct {
	nodo NodeStd
	First, Last *Node
	Count       int
}
func NewList() *List {
	return &List{nil,nil, 0}
}
func NewListEst() *ListEst {
	return &ListEst{0,nil,nil, 0}
}
func NewNode( seccion int, listEst *ListEst) *Node {
	return &Node{seccion, listEst,nil,nil}
}
func NewNodeEst( carne int) *NodeStd {
	return &NodeStd{carne,nil,nil}
}

func Insert(seccion int, list *List, listEst *ListEst) {
	var newNode = NewNode(seccion,listEst)
	if list.First == nil {
		list.First = newNode
		list.Last = newNode
		list.Count++
	} else {
		list.Last.Next = newNode
		newNode.Back = list.Last
		list.Last = newNode
		list.Count++
	}
}
func InsertEst(seccion int, list *List) {
	var newNode = NewNodeEst(seccion)
	if list.First == nil {
		list.First = newNode
		list.Last = newNode
		list.Count++
	} else {
		list.Last.Next = newNode
		newNode.Back = list.Last
		list.Last = newNode
		list.Count++
	}
}
func Delete(list *List, nameStr string) bool {
	found := false
	if list.Count > 0 {
		actual := list.First
		// it will stop when FOUND be true
		for actual != nil && !found {
			// To see if two have the same value
			found = actual.Store.Name == nameStr
			// follow with next NODE if still don't find them
			if !found {
				actual = actual.Next
			}
		}
		if found {
			if actual == list.First {
				// DELETE first node
				list.First = actual.Next
			} else {
				// DELETE node in middle or final
				actual.Back.Next = actual.Next
				if actual.Next != nil {
					actual.Next.Back = actual.Back
				}
			}
			// Decrement count of list
			list.Count--
			actual = nil
		}
	}
	return found
}
func main() {
	listSec := NewList()
	listStd := NewListEst()
	Insert()
	Insert(1,listSec,listStd)
}
