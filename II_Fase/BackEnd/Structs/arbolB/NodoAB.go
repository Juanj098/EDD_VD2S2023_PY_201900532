package arbolB

type Tutor struct {
	Carnet      int
	Name        string
	Curso       int
	Password    string
	Libros      []*Libros
	Publicacion []*Publicacion
}
type Libros struct {
	Date   string
	Name   string
	Estado string
	Tutor  int
}

type Publicacion struct {
	Tutor       int
	Publicacion string
}

type NodoB struct {
	Dato  *Tutor
	Next  *NodoB
	Prev  *NodoB
	Left  *RamaB
	Right *RamaB
}
