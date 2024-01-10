package grafos

type Grafo struct {
	Principal *NodoLista
}

func (g *Grafo) Insertar(curso string, post string) {
	if g.Principal == nil {
		g.Ifila(curso)
		g.Icolumna(curso, post)
	} else {
		g.Icolumna(curso, post)
	}
}

func (g *Grafo) Ifila(curso string) {
	nNodo := &NodoLista{Value: curso}
	if g.Principal == nil {
		g.Principal = nNodo
	} else {
		aux := g.Principal
		for aux.Abajo != nil {
			if aux.Value == curso {
				return
			}
			aux = aux.Abajo
		}
		aux.Abajo = nNodo
	}
}

func (g *Grafo) Icolumna(curso string, postr string) {
	nNodo := &NodoLista{Value: postr}
	if g.Principal != nil && curso == g.Principal.Value {
		g.Ifila(postr)
		aux := g.Principal
		for aux.Sig != nil {
			aux = aux.Sig
		}
		aux.Sig = nNodo
	} else {
		g.Ifila(curso)
		aux := g.Principal
		for aux != nil {
			if aux.Value == curso {
				break
			}
			aux = aux.Abajo
		}
		if aux != nil {
			for aux.Sig != nil {
				aux = aux.Sig
			}
			aux.Sig = nNodo
		}
	}
}
func (g *Grafo) Reporte() {
	cadena := ""
	nombre_archivo := "./grafo.dot"
	nombre_imagen := "./grafo.jpg"
	if g.Principal != nil {
		cadena += "digraph grafoDirigido{ \n rankdir=LR; \n node [shape=box]; layout=neato; \n nodo" + g.Principal.Value + "[label=\"" + g.Principal.Value + "\"]; \n"
		cadena += "node [shape = ellipse]; \n"
		cadena += g.retornarValoresMatriz()
		cadena += "\n}"
	}
	crearArchivo(nombre_archivo)
	escribirArchivo(cadena, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}

func (g *Grafo) retornarValoresMatriz() string {
	cadena := ""
	/*CREACION DE NODOS*/
	aux := g.Principal.Abajo //Filas
	aux1 := aux              //Columnas
	/*CREACION DE NODOS CON LABELS*/
	for aux != nil {
		for aux1 != nil {
			cadena += "nodo" + aux1.Value + "[label=\"" + aux1.Value + "\" ]; \n"
			aux1 = aux1.Sig
		}
		if aux != nil {
			aux = aux.Abajo
			aux1 = aux
		}
	}
	/*CONEXION DE NODOS*/
	aux = g.Principal //Filas
	aux1 = aux.Sig    //Columnas
	/*CREACION DE NODOS CON LABELS*/
	for aux != nil {
		for aux1 != nil {
			cadena += "nodo" + aux.Value + " -> "
			cadena += "nodo" + aux1.Value + "[len=1.00]; \n"
			aux1 = aux1.Sig
		}
		if aux.Abajo != nil {
			aux = aux.Abajo
			aux1 = aux.Sig
		} else {
			aux = aux.Abajo
		}
	}

	return cadena
}
