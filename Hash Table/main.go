package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"math"
	"os"
	"strconv"
)

var A = (math.Sqrt(5.0) - 1) / 2

func main() {
	table := NewTable(7, 60)
	table.insert(11, 1)
	table.insert(11, 2)
	table.insert(11, 3)
	table.insert(11, 4)
	table.insert(11, 5)
	table.insert(11, 1)
	table.insert(11, 24)
	for i := 0; i < 9; i++ {
		table.insert(18+i, i)
	}
	table.print()
}

type Node struct {
	hash  int
	value int
}
type HashTable struct {
	size       int
	charge     int
	percentage int
	array      []*Node
}

func NewTable(size int, percentage int) *HashTable {
	array := make([]*Node, size)
	return &HashTable{size, 0, percentage, array}
}
func (thisHT *HashTable) insert(new int, value int) {
	pos, repeated, inLoop := 0, false, false
	pos, repeated, inLoop = thisHT.getPosition(new, value)
	// inLoop is to verify that there's not position to insert after N iterations
	if inLoop {
		// so we magnify table and insert the new node
		thisHT.magnifyTable()
		pos, repeated, inLoop = thisHT.getPosition(new, value)
	}
	if !repeated {
		newNode := Node{new, value}
		thisHT.array[pos] = &newNode
		thisHT.charge++

		actualPercentage := (thisHT.charge * 100) / thisHT.size
		// see if we'll pass of allowed percentage that is 60%
		if actualPercentage >= thisHT.percentage {
			thisHT.magnifyTable()
		}
	}
}
func (thisHT *HashTable) magnifyTable() {
	newSize := thisHT.size
	// search next prime number to thisHt.size
loop:
	for {
		newSize += 2
		i := 1
		for i < newSize {
			i++
			if newSize%i == 0 {
				if i == newSize { //verify if newSize is prime
					break loop
				} else {
					break
				}
			}
		}
	}
	// create new array biggest
	newArray := make([]*Node, newSize)
	antique := thisHT.array
	thisHT.array = newArray
	thisHT.size = newSize
	// pass all dates at new array
	aux := 0
	for i := 0; i < len(antique); i++ {
		if antique[i] != nil {
			aux, _, _ = thisHT.getPosition(antique[i].hash, antique[i].value)
			newArray[aux] = antique[i]
		}
	}
}
func (thisHT *HashTable) getPosition(clave int, valor int) (int, bool, bool) {
	i, p, auxP, countReturn, repeated, inLoop := 0, 0, 0, 0, false, false
	m := float64(len(thisHT.array)) // size of the table
	k := float64(clave)
	p = int(math.Floor(m * (k*A - math.Floor(k*A)))) // is the getPosition where como to insert new node
	auxP = p
	// if there's a collision
	for thisHT.array[auxP] != nil && thisHT.array[auxP].value != valor {
		i++
		// new getPosition that increase quadratically
		auxP = int(math.Mod(float64(p+i*i), m)) // is the getPosition where como to insert new node
		// to resolve problem of inLoop, if doesn't find a getPosition where insert
		if auxP == p+1 {
			countReturn++
			// if this condition is TRUE, is because the collision into inLoop
			if countReturn == 2 {
				fmt.Println("in loop, i'm going to magnify table to resolve it :)")
				inLoop = true
				break
			}
		}
	}
	if thisHT.array[auxP] != nil && thisHT.array[auxP].value == valor {
		repeated = true
	}
	return auxP, repeated, inLoop
}
func (thisHT *HashTable) print() {
	data := make([][]string, thisHT.size)
	for i := 0; i < len(thisHT.array); i++ {
		tmp := make([]string, 3)
		aux := thisHT.array[i]
		if aux != nil {
			tmp[0] = strconv.Itoa(i)
			tmp[1] = strconv.Itoa(aux.hash)
			tmp[2] = strconv.Itoa(aux.value)
		} else {
			tmp[0] = "-"
			tmp[1] = "-"
			tmp[2] = "-"
		}
		data[i] = tmp
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"index", "Hash", "Valor"})
	table.SetFooter([]string{"Size", strconv.Itoa(thisHT.size), "Charge", strconv.Itoa(thisHT.charge)})
	table.AppendBulk(data)
	table.Render()
}
