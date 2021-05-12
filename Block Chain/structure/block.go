package structure

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
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

type BlockChangeJSON struct {
	Header []Block
}

func (thisL *List) SaveData() {
	var blockChain BlockChangeJSON
	aux := thisL.First
	for aux != nil {
		block := NewBlock(aux.Block.Data)
		block.Index = aux.Block.Index
		block.Date = aux.Block.Date
		block.Nonce = aux.Block.Nonce
		block.Hash = aux.Block.Hash
		block.PreviousHash = aux.Block.PreviousHash
		blockChain.Header = append(blockChain.Header, *block)
		// convert to struct of JSON
		jsonBlock, _ := json.MarshalIndent(block, "", "    ")
		// path is where como to save
		path := "Block Chain/block" + strconv.Itoa(block.Index) + ".json"
		// if file exists i'm going to remove it
		var _, err = os.Stat(path)
		if !os.IsNotExist(err) {
			err := os.Remove(path)
			if err != nil {
				log.Fatal(err)
			}
		}
		// create file
		var file, _ = os.Create(path)
		defer file.Close()
		fmt.Println("Se ha creado un archivo")
		// open file
		var myFile, err2 = os.OpenFile(path, os.O_RDWR, 0644)
		if existError(err2) {
			return
		}
		defer myFile.Close()
		// write in file
		_, err = myFile.Write(jsonBlock)
		if existError(err) {
			return
		}
		// save changes
		err = myFile.Sync()
		if existError(err) {
			return
		}
		fmt.Println("Archivo actualizado existosamente.")
		aux = aux.Next
	}

}

func existError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return err != nil
}
