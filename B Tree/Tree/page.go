package Tree

const m = 5

type Page struct {
	Claves   [m]int
	Branches [m]*Page
	count    int
}

func NewPage() *Page {
	var claves [m]int
	var branches [m]*Page
	return &Page{claves, branches, 0}
}
func FullPage(actual *Page) bool {
	return actual.count == m-1
}
