package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

// structures
type User struct{
	id int
	name string
	lastName string
	cell int
}
type Tree struct{
	root *Node
}
type Node struct{
	user User
	left *Node
	right *Node
}
// create structures
func newTree()*Tree{
	return &Tree{nil}
}
func newNode(user User)*Node{
	return &Node{user,nil,nil}
}
func newUser(id int, name string, lastName string, cell int)*User{
	return &User{id,name,lastName,cell}
}
// insert nodes in Tree
func insertTree(user *User, tree *Tree)bool{
	return insertNode(&tree.root, user)
}
func insertNode(node **Node, user *User)bool {
	if *node == nil{
		*node = newNode(*user)
		return true
	}else if user.id < (*node).user.id{
		return insertNode(&(*node).left,user)
	}else if user.id > (*node).user.id{
		return insertNode(&(*node).right,user)
	}
	return false
}
// create Graph of Tree
func createGraph(tree *Tree){
	dot := "digraph G{\nnode [shape=circle];\n"
	acum := ""
	if tree.root != nil{
		roamTree(&tree.root,&acum)
	}
	dot += acum + "\n}\n"
	// create .DOT
	path := "graph.dot"
	var _,err = os.Stat(path)
	if os.IsNotExist(err){
		var file, err = os.Create(path)
		if existError(err){
			return
		}
		defer file.Close()
		fmt.Println("Has Been Created a New File")
	}
	var file, err2 = os.OpenFile(path,os.O_RDWR,0644)
	if existError(err2){
		return
	}
	defer file.Close()
	// write in DOT
	_,err = file.WriteString(dot)
	if existError(err){
		return
	}
	// save changes
	err = file.Sync()
	if existError(err){
		return
	}
	fmt.Println("File Upgraded")
	// create Graph
	path2,_ := exec.LookPath("dot")
	cmd,_ := exec.Command(path2, "-Tpng","graph.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("graph.png",cmd,os.FileMode(mode))
}
func roamTree(actual **Node, acum *string){
	if *actual != nil{
		// get information of actual Node
		*acum += "\""+fmt.Sprint(&(*actual))+"\"[label=\" Id: "+strconv.Itoa((*actual).user.id)+"\n"
		*acum+= "Name: "+(*actual).user.name+"\n"
		*acum+= "Last Name: "+(*actual).user.lastName+"\n"
		*acum+= "Cell : "+strconv.Itoa((*actual).user.cell)+"\"];\n"
		// go to subRoot Left
		if (*actual).left != nil{
			*acum += "\""+fmt.Sprint(&(*actual))+"\" -> \""+fmt.Sprint(&(*actual).left)+"\";\n"
		}// go to subRoot Right
		if (*actual).right != nil{
			*acum += "\""+fmt.Sprint(&(*actual))+"\" -> \""+fmt.Sprint(&(*actual).right)+"\";\n"
		}
		roamTree(&(*actual).left, acum)
		roamTree(&(*actual).right, acum)
	}
}
func existError(err error)bool{
	if err != nil{
		fmt.Println(err.Error())
	}
	return err != nil
}
func main (){
	user1 := newUser(4,"carlos", "castro", 54637862)
	user2 := newUser(2,"javier", "castro", 54283876)
	user3 := newUser(6,"jose", "paredes", 54283876)
	user4 := newUser(10,"mario", "lopez", 54283876)
	user5 := newUser(1,"maria", "estrada", 54283876)
	user6 := newUser(5,"maria", "estrada", 54283876)
	tree := newTree()
	insertTree(user1, tree)
	insertTree(user2, tree)
	insertTree(user3, tree)
	insertTree(user4, tree)
	insertTree(user5, tree)
	insertTree(user6, tree)
	createGraph(tree)
}