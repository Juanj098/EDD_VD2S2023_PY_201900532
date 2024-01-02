package arbolB

type Tutor struct {
	Carnet   int
	Name     string
	Curso    string
	Password string
}

type NodoB struct {
	Dato  *Tutor
	Next  *NodoB
	Prev  *NodoB
	Left  *RamaB
	Right *RamaB
}
