package arbolB

type RamaB struct {
	First    *NodoB
	Hoja     bool
	Contador int
}

func (r *RamaB) Insert(newNodo *NodoB) {
	if r.First == nil {
		r.First = newNodo
		r.Contador++
	} else {
		if newNodo.Dato.Carnet < r.First.Dato.Carnet {
			newNodo.Next = r.First
			r.First.Left = newNodo.Right
			r.First.Prev = newNodo
			r.First = newNodo
			r.Contador++
		} else if r.First.Next != nil {
			if r.First.Next.Dato.Carnet > newNodo.Dato.Carnet {
				newNodo.Next = r.First.Next
				newNodo.Prev = r.First
				r.First.Next.Left = newNodo.Right
				r.First.Right = newNodo.Left
				r.First.Next.Prev = newNodo
				r.First.Next = newNodo
				r.Contador++
			} else {
				aux := r.First.Next
				newNodo.Prev = aux
				aux.Right = newNodo.Left
				aux.Next = newNodo
				r.Contador++
			}
		} else if r.First.Next == nil {
			newNodo.Prev = r.First
			r.First.Right = newNodo.Left
			r.First.Next = newNodo
			r.Contador++
		}
	}
}
