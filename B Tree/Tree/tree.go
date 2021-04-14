package Tree

import "fmt"

type Tree struct {
	root *Page
}

func NewTree() *Tree {
	return &Tree{nil}
}
func (this *Tree) Insert(valor int) {
	insert(&this.root, valor)
}
func insert(root **Page, value int) {
	subeArriba := false
	median := 0
	var nd *Page
	Empujar(*root, value, &subeArriba, &median, &nd)
	if subeArriba { //si se produjo una reorganizacion de nodos lo cual ser dividio la raiz entonces, la bandera sube_arriba lo indica
		p := NewPage()
		p.count = 1
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
	(*nuevo).count = (orden - 1) - posMedian /* numero de claves en el nuevo nodo*/
	actual.count = posMedian                 // numero claves en el nodo origen

	/* Es insertada la clave y rama en el nodo que le corresponde*/
	if k <= orden/2 { //si k es menor al minimo de claves  que puede haber en la pagina
		PushLeaf(actual, value, rd, k) //inserta en el nodo origen
	} else {
		PushLeaf(*nuevo, value, rd, k-posMedian) //se inserta el nuevo value que traiamos en el nodo nuevo
	}

	/* se extrae clave mediana del nodo origen*/
	*median = actual.Claves[actual.count]

	/* Rama0 del nuevo nodo es la rama de la mediana*/
	(*nuevo).Branches[0] = actual.Branches[actual.count]
	actual.count--
}
func PushLeaf(actual *Page, valor int, rd *Page, k int) {
	/* desplza a la derecha los elementos para hcer un hueco*/
	for i := actual.count; i >= k+1; i-- {
		actual.Claves[i+1] = actual.Claves[i]
		actual.Branches[i+1] = actual.Branches[i]
	}
	actual.Claves[k+1] = valor
	actual.Branches[k+1] = rd
	actual.count++
}
func SearchNode(actual *Page, valor int, k *int) bool {
	/*tomar en cuenta que k es la direccion de las ramas por las que puede bajar la busqueda*/
	var found bool
	if valor < actual.Claves[1] { //ese 1 significa que busca desde la primera posicion en claves por tanto si cumple el valor se va a los valores menores
		*k = 0 //nos indica que bajaresmo por la rama 0
		found = false
	} else //examina las claves del nodo en orden descendente
	{
		*k = actual.count                             //desde la clave actual
		for (valor < actual.Claves[*k]) && (*k > 1) { //buscar una posicion hasta donde valor deje de ser menor ( por si vienen un valor menor a los que hay en el nodo )
			*k--
		}
		found = valor == actual.Claves[*k] //si la posicion encontrada es igual al valor ; lo cual clave repetida
	}
	return found
}
