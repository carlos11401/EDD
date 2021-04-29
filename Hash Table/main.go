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
	table.insert(100, 25)
	table.insert(101, 25)
	table.insert(35, 25)
	table.insert(2, 25)
	table.insert(1, 25)
	table.insert(18, 25)
	for i := 0; i < 11; i++ {
		table.insert(18, i)
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
	newNode := Node{new, value}
	pos := thisHT.getPosition(new, value)
	thisHT.array[pos] = &newNode
	thisHT.charge++
	actualPercentage := (thisHT.charge * 100) / thisHT.size
	// see if we'll pass of allowed percentage that is 60%
	if actualPercentage >= thisHT.percentage {
		sizeNuevo := thisHT.size
		foundSize := false
		// search next prime number to thisHt.size
		for !foundSize {
			sizeNuevo += 2
			i := 1
			for i < sizeNuevo {
				i++
				mod := sizeNuevo % i
				if mod == 0 {
					break
				}
			}
			if i == sizeNuevo {
				foundSize = true
			}
		}
		// create new array biggest
		newArray := make([]*Node, sizeNuevo)
		antique := thisHT.array
		thisHT.array = newArray
		thisHT.size = sizeNuevo
		// pass all dates at new array
		aux := 0
		for i := 0; i < len(antique); i++ {
			if antique[i] != nil {
				aux = thisHT.getPosition(antique[i].hash, antique[i].value)
				newArray[aux] = antique[i]
			}
		}
	}
}
func (thisHT *HashTable) getPosition(clave int, valor int) int {
	i, p, auxP, countReturn := 0, 0, 0, 0
	m := float64(len(thisHT.array)) // size of the table
	k := float64(clave)
	p = int(math.Floor(m * (k*A - math.Floor(k*A)))) // is the getPosition where como to insert new node
	auxP = p
	// if there's a collision
	for thisHT.array[auxP] != nil && thisHT.array[auxP].value != valor {
		i++
		// new getPosition that increase quadratically
		auxP = p + i*i
		// see if i can insert after of getPosition (P) because i could pass other getPosition of the array major than len(Array)
		if auxP > (len(thisHT.array) - 1) {
			// calculate the value nearest to auxP
			auxPivot := int(math.Floor(float64(auxP / len(thisHT.array)))) // approach down
			// nearestP = auxPivot
			newP := auxP - auxPivot*len(thisHT.array) - p
			// see if i can insert after of getPosition (P) because i could pass other getPosition of the array major than len(Array)
			if (p + newP) < len(thisHT.array) {
				auxP = p + newP
			} else {
				aux := len(thisHT.array) - 1 - p
				auxP = newP - aux - 1
			}
		}
		// to resolve problem of loop, if doesn't find a getPosition where insert
		if auxP == p+1 {
			countReturn++
			// if this condition is TRUE, is because the collision into loop
			if countReturn == 2 {
				fmt.Println("there's not space where add new node, try change a key")
				break
			}
		}
	}
	return auxP
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
