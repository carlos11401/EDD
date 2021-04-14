package Tree

const m = 5

type Page struct {
	Claves   [m]int
	Branches [m]*Page
	Count    int
}

func NewPage() *Page {
	var claves [m]int
	var branches [m]*Page
	return &Page{claves, branches, 0}
}
func FullPage(actual *Page) bool {
	return actual.Count == m-1
}
