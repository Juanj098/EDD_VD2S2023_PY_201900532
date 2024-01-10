package grafos

type NodoLista struct {
	Sig   *NodoLista
	Abajo *NodoLista
	Value string
}
