package THash

type Estudiante struct {
	Carnet   int
	Name     string
	Password string
	// Cursos   []string
}

type NodoHash struct {
	Llave int
	Dato  *Estudiante
}
