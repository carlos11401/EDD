package tree

import (
	"container/list"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"math"
)

type Node struct {
	value string
	right *Node
	left  *Node
}
type Tree struct {
	root *Node
}

func (thisNode *Node) sum() string {
	if thisNode.right != nil && thisNode.left != nil {
		return thisNode.right.value + thisNode.left.value
	}
	return "-1"
}
func newNode(value string, right *Node, left *Node) *Node {
	return &Node{value, right, left}
}
func NewTree() *Tree {
	return &Tree{}
}
func (thisTree *Tree) Insert(value string) {
	valueEncrypted := Encrypt(value)
	n := newNode(valueEncrypted, nil, nil)
	if thisTree.root == nil {
		listNodes := list.New()
		listNodes.PushBack(n)
		listNodes.PushBack(newNode(Encrypt("-1"), nil, nil))
		thisTree.buildTree(listNodes)
	} else {
		listNodes := thisTree.GetList()
		listNodes.PushBack(n)
		thisTree.buildTree(listNodes)
	}
}
func (thisTree *Tree) GetList() *list.List {
	listNodes := list.New()
	getList(listNodes, thisTree.root.left)
	getList(listNodes, thisTree.root.right)
	return listNodes
}

// getting last nodes in tree ROAM(in_order)
func getList(listNodes *list.List, actual *Node) {
	if actual != nil {
		getList(listNodes, actual.left)
		// verify if is the last node
		if actual.right == nil && actual.value != Encrypt("-1") {
			listNodes.PushBack(actual)
		}
		getList(listNodes, actual.right)
	}
}
func (thisTree *Tree) buildTree(listNodes *list.List) {
	size := float64(listNodes.Len())
	noLevels := 1
	// it will calculate the number of levels that tree has
	for (size / 2) > 1 {
		noLevels++
		size = size / 2
	}
	// nodesTot is the number of total nodes in base to number of levels
	nodesTot := math.Pow(2, float64(noLevels))
	// insert -1 to fill the list
	for listNodes.Len() < int(nodesTot) {
		listNodes.PushBack(newNode(Encrypt("-1"), nil, nil))
	}
	// reducing list in half
	for listNodes.Len() > 1 {
		first := listNodes.Front()
		second := first.Next()
		listNodes.Remove(first)
		listNodes.Remove(second)
		// casting value of list because return a interface
		node1 := first.Value.(*Node)
		node2 := second.Value.(*Node)
		nuevo := newNode(node1.value+node2.value, node2, node1)
		// add new node2, here is where list come to reducing in half
		listNodes.PushBack(nuevo)
	}
	// how list just have one element we know that is the root of out tree
	thisTree.root = listNodes.Front().Value.(*Node)
}

var key = []byte("asdfasdfasdfasdfasdfasdfasdfasdf")

func Encrypt(date string) string {
	plaintext := []byte(date)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	nonce := []byte("gopostmedium")
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}
func Decrypt(date string) string {
	ciphertext, _ := hex.DecodeString(date)
	nonce := []byte("gopostmedium")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return fmt.Sprintf("%s", plaintext)
}
