package Tree

import "fmt"

type Tree struct {
	Root *Page
}

func NewTree() *Tree {
	return &Tree{nil}
}

// to insert
func (this *Tree) Insert(valor int) {
	insert(&this.Root, valor)
}
func insert(root **Page, value int) {
	subeArriba := false
	median := 0
	var nd *Page
	Empujar(*root, value, &subeArriba, &median, &nd)
	if subeArriba { //si se produjo una reorganizacion de nodos lo cual ser dividio la raiz entonces, la bandera sube_arriba lo indica
		p := NewPage()
		p.Count = 1
		p.Claves[1] = median
		p.Branches[0] = *root
		p.Branches[1] = nd
		*root = p
	}
}
func Empujar(actual *Page, valor int, subeArriba *bool, median *int, nuevo **Page) {
	k := 0
	if actual == nil {
		*subeArriba = true
		*median = valor
		nuevo = nil
	} else {
		var is bool
		is = SearchNode(actual, valor, &k) //k busca la rama por eso se pasa por referencia
		if is {
			fmt.Println("Clave Duplicada: ", valor)
			*subeArriba = false
			return
		}
		Empujar(actual.Branches[k], valor, subeArriba, median, nuevo)
		/* devuelve control vuelve por el camino de busqueda*/
		if *subeArriba {
			if FullPage(actual) {
				DivideNode(actual, *median, *nuevo, k, median, nuevo)
			} else {
				*subeArriba = false
				PushLeaf(actual, *median, *nuevo, k)
			}
		}
	}
}
func DivideNode(actual *Page, value int, rd *Page, k int, median *int, nuevo **Page) {
	posMedian := 3
	orden := 5
	*nuevo = NewPage() //se crea una nueva pagina
	for i := posMedian + 1; i < orden; i++ {
		/* es desplzada la mida derecha al nuevo nodo, la clave mediana se queda en el nodo origen*/
		(*nuevo).Claves[i-posMedian] = actual.Claves[i]
		(*nuevo).Branches[i-posMedian] = actual.Branches[i]
	}
	(*nuevo).Count = (orden - 1) - posMedian /* numero de claves en el nuevo nodo*/
	actual.Count = posMedian                 // numero claves en el nodo origen

	/* Es insertada la clave y rama en el nodo que le corresponde*/
	if k <= orden/2 { //si k es menor al minimo de claves  que puede haber en la pagina
		PushLeaf(actual, value, rd, k) //inserta en el nodo origen
	} else {
		PushLeaf(*nuevo, value, rd, k-posMedian) //se inserta el nuevo value que traiamos en el nodo nuevo
	}

	/* se extrae clave mediana del nodo origen*/
	*median = actual.Claves[actual.Count]

	/* Rama0 del nuevo nodo es la rama de la mediana*/
	(*nuevo).Branches[0] = actual.Branches[actual.Count]
	actual.Count--
}
func PushLeaf(actual *Page, valor int, rd *Page, k int) {
	/* desplza a la derecha los elementos para hcer un hueco*/
	for i := actual.Count; i >= k+1; i-- {
		actual.Claves[i+1] = actual.Claves[i]
		actual.Branches[i+1] = actual.Branches[i]
	}
	actual.Claves[k+1] = valor
	actual.Branches[k+1] = rd
	actual.Count++
}
func SearchNode(actual *Page, valor int, k *int) bool {
	/*tomar en cuenta que k es la direccion de las ramas por las que puede bajar la busqueda*/
	var found bool
	if valor < actual.Claves[1] { //ese 1 significa que busca desde la primera posicion en claves por tanto si cumple el valor se va a los valores menores
		*k = 0 //nos indica que bajaresmo por la rama 0
		found = false
	} else //examina las claves del nodo en orden descendente
	{
		*k = actual.Count                             //desde la clave actual
		for (valor < actual.Claves[*k]) && (*k > 1) { //buscar una posicion hasta donde valor deje de ser menor ( por si vienen un valor menor a los que hay en el nodo )
			*k--
		}
		found = valor == actual.Claves[*k] //si la posicion encontrada es igual al valor ; lo cual clave repetida
	}
	return found
}

// to delete
func (this *Tree) Delete(valor int) {
	delete(&this.Root, valor)
}
func delete(root **Page, value int) {
	found := false
	deleteRegister(*root, value, &found)
	if found {
		if (*root).Count == 0 {
			/* la raiz esta vacia, se libera el nodo*/
			p := NewPage()
			*p = **root
			*root = (*root).Branches[0]
		}
	} // else la clave no esta en le arbol
}
func deleteRegister(actual *Page, value int, found *bool) {
	var k int
	if actual != nil {
		*found = SearchNode(actual, value, &k)
		if *found {
			if actual.Branches[k-1] == nil {
				quitar(actual, k)
			} else {
				sucesor(actual, k)
				deleteRegister(actual.Branches[k], actual.Claves[k], found)
			}
		} else {
			deleteRegister(actual.Branches[k], value, found)
		}
		if actual.Branches[k] != nil {
			if actual.Branches[k].Count < m/2 {
				restablecer(actual, k)
			}
		}
	} else {
		*found = false
	}
}
func restablecer(actual *Page, k int) {
	if k > 0 {
		if actual.Branches[k-1].Count > m/2 {
			moveRight(actual, k)
		} else {
			combine(actual, k)
		}
	} else {
		if actual.Branches[1].Count > m/2 {
			moveLeft(actual, 1)
		} else {
			combine(actual, 1)
		}
	}
}
func combine(actual *Page, k int) {
	leftNode := NewPage()
	q := NewPage()

	q = actual.Branches[k]
	leftNode = actual.Branches[k-1]

	leftNode.Count++
	leftNode.Claves[leftNode.Count] = actual.Claves[k]
	leftNode.Branches[leftNode.Count] = q.Branches[0]

	for j := 1; j <= q.Count; j++ {
		leftNode.Count++
		leftNode.Claves[leftNode.Count] = q.Claves[j]
		leftNode.Branches[leftNode.Count] = q.Branches[j]
	}
	for j := k; j <= actual.Count-1; j++ {
		actual.Claves[j] = actual.Claves[j+1]
		actual.Branches[j] = actual.Branches[j+1]
	}
	actual.Count--
}
func moveLeft(actual *Page, k int) {
	problemNode := NewPage()
	rightNode := NewPage()
	problemNode = actual.Branches[k-1]
	rightNode = actual.Branches[k]

	problemNode.Count++
	problemNode.Claves[problemNode.Count] = actual.Claves[k]
	problemNode.Branches[problemNode.Count] = rightNode.Branches[0]

	actual.Claves[k] = rightNode.Claves[1]
	rightNode.Branches[1] = rightNode.Branches[0]
	rightNode.Count--

	for i := 1; i <= rightNode.Count; i++ {
		rightNode.Claves[i] = rightNode.Claves[i+1]
		rightNode.Branches[i] = rightNode.Branches[i+1]
	}
}
func moveRight(actual *Page, k int) {
	problemNode := NewPage()
	leftNode := NewPage()
	problemNode = actual.Branches[k]
	leftNode = actual.Branches[k-1]
	for j := problemNode.Count; j >= 1; j-- {
		problemNode.Claves[j+1] = problemNode.Claves[j]
		problemNode.Branches[j+1] = problemNode.Branches[j]
	}
	problemNode.Count++
	problemNode.Branches[1] = problemNode.Branches[0]

	problemNode.Claves[1] = actual.Claves[k]

	actual.Claves[k] = leftNode.Claves[leftNode.Count]
	problemNode.Branches[0] = leftNode.Branches[leftNode.Count]
	leftNode.Count--
}
func sucesor(actual *Page, k int) {
	q := NewPage()
	q = actual.Branches[k]
	for q.Branches[0] != nil {
		q = q.Branches[0]
	}
	actual.Claves[k] = q.Claves[1]
}
func quitar(actual *Page, k int) {
	for j := k + 1; j <= actual.Count; j++ {
		actual.Claves[j-1] = actual.Claves[j]
		actual.Branches[j-1] = actual.Branches[j]
	}
	actual.Count--
}
