package THash

type Estudiante struct {
	Carnet   int
	Name     string
	Password string
	Cursos   []int
}

type NodoHash struct {
	Llave int
	Dato  *Estudiante
}
