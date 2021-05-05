package structure

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Node struct {
	Block *Block
	Next  *Node
	Back  *Node
}
type List struct {
	First, Last *Node
	Count       int
}
type Block struct {
	Index        int
	Date         string
	Data         string
	Nonce        int
	PreviousHash string
	Hash         string
}

func NewBlock(data string) *Block {
	return &Block{0, "", data, 0, "", ""}
}
func NewList() *List {
	return &List{nil, nil, 0}
}
func NewNode(block *Block) *Node {
	return &Node{block, nil, nil}
}

// Insert func for vector
func (thisL *List) Insert(data string) {
	newBlock := NewBlock(data)
	// get index in base at the count of this list
	newBlock.Index, newBlock.PreviousHash = thisL.GetIndexAndPreHash()
	// to get time of buy
	t := time.Now()
	newBlock.Date = fmt.Sprintf("%d-%02d-%02d::%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	// to get nonce and hash
	newBlock.MineBlock()
	var newNode = NewNode(newBlock)
	if thisL.First == nil {
		thisL.First = newNode
		thisL.Last = newNode
		thisL.Count++
	} else {
		thisL.Last.Next = newNode
		newNode.Back = thisL.Last
		thisL.Last = newNode
		thisL.Count++
	}
}
func (thisL *List) GetIndexAndPreHash() (int, string) {
	if thisL.Count == 0 {
		return thisL.Count, "0000"
	} else {
		return thisL.Count, thisL.Last.Block.Hash
	}
}
func (thisL *List) Print() {
	aux := thisL.First
	for aux != nil {
		fmt.Println(strconv.Itoa(aux.Block.Index) + ".")
		fmt.Println("     > PreHash: " + aux.Block.PreviousHash)
		fmt.Println("     > Hash: " + aux.Block.Hash)
		fmt.Println("     > Nonce: " + strconv.Itoa(aux.Block.Nonce))
		fmt.Println("     > Date: " + aux.Block.Date)
		fmt.Println("     > Data: " + aux.Block.Data)
		aux = aux.Next
	}
}
func (thisB *Block) GetSha256() string {
	var str string
	str = strconv.Itoa(thisB.Index) +
		thisB.Date +
		thisB.PreviousHash +
		thisB.Data +
		strconv.Itoa(thisB.Nonce)
	hash := sha256.New()
	hash.Write([]byte(str))
	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}
func (thisB *Block) MineBlock() {
	for i := 0; ; i++ {
		thisB.Nonce = i
		str := thisB.GetSha256()
		if strings.HasPrefix(str, "0000") {
			thisB.Hash = str
			break
		}
	}
}
