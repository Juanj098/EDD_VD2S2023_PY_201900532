package arbolB

import (
	"strconv"
)

type ArbolB struct {
	Raiz  *RamaB
	Orden int
}

func (a *ArbolB) InsertaRama(nodo *NodoB, rama *RamaB) *NodoB {
	if rama.Hoja {
		rama.Insert(nodo)
		if rama.Contador == a.Orden {
			return a.Dividir(rama)
		} else {
			return nil
		}
	} else {
		temp := rama.First
		for temp != nil {
			if nodo.Dato == temp.Dato { //comparar y no guardar si son iguales
				return nil
			} else if nodo.Dato.Carnet < temp.Dato.Carnet {
				obj := a.InsertaRama(nodo, temp.Left)
				if obj != nil {
					rama.Insert(obj)
					if rama.Contador == a.Orden {
						return a.Dividir(rama)
					}
				}
				return nil
			} else if temp.Next == nil {
				obj := a.InsertaRama(nodo, temp.Right)
				if obj != nil {
					rama.Insert(obj)
					if rama.Contador == a.Orden {
						return a.Dividir(rama)
					}
				}
				return nil
			}
			temp = temp.Next
		}
	}
	return nil
}

func (a *ArbolB) Dividir(rama *RamaB) *NodoB {
	val := &NodoB{Dato: &Tutor{Carnet: 000000000, Name: "", Curso: "", Password: ""}}
	aux := rama.First
	rDerecha := &RamaB{First: nil, Contador: 0, Hoja: true}
	rIzquierda := &RamaB{First: nil, Contador: 0, Hoja: true}
	cTemp := 0
	for aux != nil {
		cTemp++
		if cTemp < 2 {
			temp := &NodoB{Dato: aux.Dato}
			temp.Left = aux.Left
			if cTemp == 1 {
				temp.Right = aux.Next.Left
			}
			if temp.Right != nil && temp.Left != nil {
				rIzquierda.Hoja = false
			}
			rIzquierda.Insert(temp)
		} else if cTemp == 2 {
			val.Dato = aux.Dato
		} else {
			temp := &NodoB{Dato: aux.Dato}
			temp.Left = aux.Left
			temp.Right = aux.Right
			if temp.Right != nil && temp.Left != nil {
				rDerecha.Hoja = false
			}
			rDerecha.Insert(temp)
		}
		aux = aux.Next
	}
	nuevoN := &NodoB{Dato: val.Dato}
	nuevoN.Right = rDerecha
	nuevoN.Left = rIzquierda
	return nuevoN
}

func (a *ArbolB) InsertA(carnet int, name string, curso string, pass string) {
	newNodo := &NodoB{Dato: &Tutor{Carnet: carnet, Name: name, Curso: curso, Password: pass}}
	if a.Raiz == nil {
		a.Raiz = &RamaB{First: nil, Hoja: true, Contador: 0}
		a.Raiz.Insert(newNodo)
	} else {
		// funciones insertar
		obj := a.InsertaRama(newNodo, a.Raiz)
		if obj != nil {
			a.Raiz = &RamaB{First: nil, Contador: 0, Hoja: true}
			a.Raiz.Insert(obj)
			a.Raiz.Hoja = false
		}
	}
}

func (a *ArbolB) Buscar(carnet string, listaSimple *ListaSimple) *NodoB {
	temp, _ := strconv.Atoi(carnet)
	B := listaSimple.Buscar(temp)
	if B != nil {
		return B.Tutor
	}
	if a.Raiz != nil && a.Raiz.First != nil {
		a.BuscarTree(a.Raiz.First, temp, listaSimple)
	}
	B = listaSimple.Buscar(temp)
	if B != nil {
		return B.Tutor
	}
	return nil
}

func (a *ArbolB) BuscarTree(raiz *NodoB, carnet int, listaSimpe *ListaSimple) {
	if raiz != nil {
		aux := raiz
		for aux != nil {
			if aux.Left != nil {
				a.BuscarTree(aux.Left.First, carnet, listaSimpe)
			}
			if aux.Dato.Carnet == carnet {
				listaSimpe.Insertar(aux)
			}
			if aux.Next == nil {
				if aux.Right != nil {
					a.BuscarTree(aux.Right.First, carnet, listaSimpe)
				}
			}
			aux = aux.Next
		}
	}
}

// Crear Grafo...
func (a *ArbolB) Graficar() {
	nameDot := "./ArbolB.dot"
	nameImg := "./ArbolB.jpg"
	cadena := `digraph arbol { 
		node[shape=record]`
	if a.Raiz != nil {
		cadena += a.grafo(a.Raiz.First)
		cadena += a.conexiones(a.Raiz.First)
		cadena += "}"
	}
	crearArchivo(cadena)
	escribirArchivo(cadena, nameDot)
	ejecutar(nameImg, nameDot)
}

func (a *ArbolB) grafo(rama *NodoB) string {
	var cadena string
	if rama != nil {
		cadena += a.grafoR(rama)
		aux := rama
		for aux != nil {
			if aux.Left != nil {
				cadena += a.grafo(aux.Left.First)
			}
			if aux.Next == nil {
				if aux.Right != nil {
					cadena += a.grafo(aux.Right.First)
				}
			}
			aux = aux.Next
		}
	}
	return cadena

}

func (a *ArbolB) grafoR(rama *NodoB) string {
	var cadena string
	if rama != nil {
		aux := rama
		cadena += cadena + "R" + strconv.Itoa(rama.Dato.Carnet) + "[label=\""
		r := 1
		for aux != nil {
			if aux.Left != nil {
				cadena += cadena + "<C" + strconv.Itoa(r) + ">|"
				r++
			}
			if aux.Next != nil {
				cadena += cadena + strconv.Itoa(aux.Dato.Carnet) + "|"
			} else {
				cadena += cadena + strconv.Itoa(aux.Dato.Carnet)
				if aux.Right != nil {
					cadena += cadena + "|<C" + strconv.Itoa(r) + ">"
				}
			}
			aux = aux.Next
		}
		cadena += cadena + "\"];\n"
	}
	return cadena
}

func (a *ArbolB) conexiones(rama *NodoB) string {
	var cadena string
	cadena = ""
	if rama != nil {
		aux := rama
		current := "R" + strconv.Itoa(rama.Dato.Carnet)
		r := 1
		for aux != nil {
			if aux.Left != nil {
				cadena += current + ":C" + strconv.Itoa(r) + " -> " + "R" + strconv.Itoa(aux.Left.First.Dato.Carnet) + ";\n"
				r++
				cadena += a.conexiones(aux.Left.First)
			}
			if aux.Next != nil {
				if aux.Right != nil {
					cadena += current + ":C" + strconv.Itoa(r) + " -> " + "R" + strconv.Itoa(aux.Right.First.Dato.Carnet) + ";\n"
					r++
					cadena += a.conexiones(aux.Right.First)
				}
			}
			aux = aux.Next
		}

	}
	return cadena
}
