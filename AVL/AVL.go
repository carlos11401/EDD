package main

// structures of AVL
type Node struct {
	Value  int
	Factor int
	Left   *Node
	Right  *Node
}
type Tree struct {
	root *Node
}

// func to create new structures
func NewTree() *Tree {
	return &Tree{nil}
}
func NewNode(value int) *Node {
	return &Node{value, 0, nil, nil}
}

// to rotate nodes
func rotationII(n *Node, n1 *Node) *Node {
	n.Left = n1.Right
	n1.Right = n
	if n1.Factor == -1 {
		n.Factor = 0
		n1.Factor = 0
	} else {
		n.Factor = -1
		n1.Factor = 1
	}
	return n1
}
func rotationDD(n *Node, n1 *Node) *Node {
	n.Right = n1.Left
	n1.Left = n
	if n1.Factor == 1 {
		n.Factor = 0
		n1.Factor = 0
	} else {
		n.Factor = 1
		n1.Factor = -1
	}
	return n1
}
func rotationDI(n *Node, n1 *Node) *Node {
	n2 := n1.Left
	n.Right = n2.Left
	n2.Left = n
	n1.Left = n2.Right
	n2.Right = n1
	if n2.Factor == 1 {
		n.Factor = -1
	} else {
		n.Factor = 0
	}
	if n2.Factor == -1 {
		n1.Factor = 1
	} else {
		n1.Factor = 0
	}
	n2.Factor = 0
	return n2
}
func rotationID(n *Node, n1 *Node) *Node {
	n2 := n1.Right
	n.Left = n2.Right
	n2.Right = n
	n1.Right = n2.Left
	n2.Left = n1
	if n2.Factor == 1 {
		n1.Factor = -1
	} else {
		n1.Factor = 0
	}
	if n2.Factor == -1 {
		n.Factor = 1
	} else {
		n.Factor = 0
	}
	n2.Factor = 0
	return n2
}

// to insert nodes and rotate
func insert(root *Node, value int, hc *bool) *Node {
	var n1 *Node
	if root == nil {
		root = NewNode(value)
		*hc = true
	} else if value < root.Value {
		left := insert(root.Left, value, hc)
		root.Left = left
		if *hc {
			switch root.Factor {
			case 1:
				root.Factor = 0
				*hc = false
				break
			case 0:
				root.Factor = -1
				break
			case -1:
				n1 = root.Left
				if n1.Factor == -1 {
					root = rotationII(root, n1)
				} else {
					root = rotationID(root, n1)
				}
				*hc = false
			}
		}
	} else if value > root.Value {
		right := insert(root.Right, value, hc)
		root.Right = right
		if *hc {
			switch root.Factor {
			case 1:
				n1 = root.Right
				if n1.Factor == 1 {
					root = rotationDD(root, n1)
				} else {
					root = rotationDI(root, n1)
				}
				*hc = false
				break
			case 0:
				root.Factor = 1
				break
			case -1:
				root.Factor = 0
				*hc = false
			}

		}
	}
	return root
}

func (this *Tree) Insert(valor int) {
	b := false
	a := &b
	this.root = insert(this.root, valor, a)
}
