package arbolB

type NodoLista struct {
	Tutor     *NodoB
	Siguiente *NodoLista
}

type ListaSimple struct {
	Inicio   *NodoLista
	Longitud int
}

func (l *ListaSimple) Insertar(tutor *NodoB) {
	if l.Buscar(tutor.Dato.Carnet) != nil {
		return
	}
	if l.Longitud == 0 {
		nuevo := &NodoLista{Tutor: tutor, Siguiente: nil}
		l.Inicio = nuevo
		l.Longitud++
	} else {
		aux := l.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = &NodoLista{Tutor: tutor, Siguiente: nil}
		l.Longitud++
	}
}

func (l *ListaSimple) Buscar(Carnet int) *NodoLista {
	aux := l.Inicio
	for aux != nil {
		if aux.Tutor.Dato.Carnet == Carnet {
			return aux
		}
		aux = aux.Siguiente
	}
	return nil
}
