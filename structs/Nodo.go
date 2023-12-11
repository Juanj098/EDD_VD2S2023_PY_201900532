package structs

type Alumnos struct {
	Name   string
	Carnet int
}

type NodoLD struct {
	Alumno *Alumnos
	Next   *NodoLD
	Prev   *NodoLD
}

type Tutor struct {
	Carnet int
	Name   string
	Curso  string
	Prom   int
}

type NodoLDC struct {
	Tutor *Tutor
	Next  *NodoLDC
	Prev  *NodoLDC
}
