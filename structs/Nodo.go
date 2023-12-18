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

type NodoCola struct {
	Tutor     *Tutor
	Sig       *NodoCola
	Prioridad int
}

type Data struct {
	Ctutor      int
	Cestudiante int
	Curso       string
}

type NodoMtx struct {
	Next *NodoMtx
	Prev *NodoMtx
	Up   *NodoMtx
	PosX int
	PosY int
	Down *NodoMtx
	Dato *Data
}

type NodoTree struct {
	Left         *NodoTree
	Right        *NodoTree
	Value        string
	Altura       int
	F_equilibrio int
}
