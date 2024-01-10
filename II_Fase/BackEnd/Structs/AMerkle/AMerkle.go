package AMerkle

import (
	"encoding/hex"
	"math"
	"strconv"
	"time"

	"golang.org/x/crypto/sha3"
)

type ArbolMerkle struct {
	Root  *NodoMerkle
	BData *NBloque
	Cant  int
}

func Date() string {
	now := time.Now()
	formato := "02-01-2006::15:04:30"
	Date := now.Format(formato)
	return Date
}

func (a *ArbolMerkle) AgregarB(state string, name string, tutor int) {
	nRegistro := &InfoBloque{Date: Date(), State: state, Name: name, Tutor: tutor}
	nBloque := &NBloque{Value: nRegistro}
	if a.BData == nil {
		a.BData = nBloque
		a.Cant++
	} else {
		aux := a.BData
		for aux.Sig != nil {
			aux = aux.Sig
		}
		nBloque.Ant = aux
		aux.Sig = nBloque
		a.Cant++
	}
}

func (a *ArbolMerkle) Crypto(cadena string) string {
	hash := sha3.New256()
	hash.Write([]byte(cadena))
	encrypt := hex.EncodeToString(hash.Sum(nil))
	return encrypt
}

func (a *ArbolMerkle) GenArbol() {
	nvl := 1
	for int(math.Pow(2, float64(nvl))) < a.Cant {
		nvl++
	}
	for i := a.Cant; i < int(math.Pow(2, float64(nvl))); i++ {
		a.AgregarB("nulo", "nulo", 0)
	}
	a.Ghash()
}

func (a *ArbolMerkle) Ghash() {
	var arrN []*NodoMerkle
	aux := a.BData
	for aux != nil {
		cadena := aux.Value.Date + aux.Value.State + aux.Value.Name + strconv.Itoa(aux.Value.Tutor)
		encrypto := a.Crypto(cadena)
		nHoja := &NodoMerkle{Valor: encrypto, B_info: aux}
		arrN = append(arrN, nHoja)
		aux = aux.Sig
	}
	a.Root = a.CrearArbol(arrN)
}

func (a *ArbolMerkle) CrearArbol(arrN []*NodoMerkle) *NodoMerkle {
	var aux []*NodoMerkle
	var raiz *NodoMerkle
	if len(aux) == 2 {
		encriptado := a.Crypto(arrN[0].Valor + arrN[1].Valor)
		raiz = &NodoMerkle{Valor: encriptado}
		raiz.Left = arrN[0]
		raiz.Right = arrN[1]
		return raiz
	} else {
		for i := 0; i < len(arrN); i += 2 {
			encriptacion := a.Crypto(arrN[i].Valor + arrN[i+1].Valor)
			nRaiz := &NodoMerkle{Valor: encriptacion}
			nRaiz.Left = arrN[i]
			nRaiz.Right = arrN[i+1]
			aux = append(arrN, nRaiz)
		}
		return a.CrearArbol(aux)
	}
}

func (a *ArbolMerkle) Graficar() {
	cadena := ""
	nombre_archivo := "./arbolMerkle.dot"
	nombre_imagen := "./arbolMerkle.jpg"
	if a.Root != nil {
		cadena += "digraph arbol { node [shape=box];"
		cadena += a.retornarValoresArbol(a.Root, 0)
		cadena += "}"
	}
	crearArchivo(nombre_archivo)
	escribirArchivo(cadena, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}

func (a *ArbolMerkle) retornarValoresArbol(raiz *NodoMerkle, indice int) string {
	cadena := ""
	numero := indice + 1
	if raiz != nil {
		cadena += "\""
		cadena += raiz.Valor[:20]
		cadena += "\" [dir=back];\n"
		if raiz.Left != nil && raiz.Right != nil {
			cadena += "\""
			cadena += raiz.Valor[:20]
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Left, numero)
			cadena += "\""
			cadena += raiz.Valor[:20]
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Right, numero)
			cadena += "{rank=same" + "\"" + (raiz.Left.Valor[:20]) + "\"" + " -> " + "\"" + (raiz.Right.Valor[:20]) + "\"" + " [style=invis]}; \n"
		}
	}
	if raiz.B_info != nil {
		cadena += "\""
		cadena += raiz.Valor[:20]
		cadena += "\" -> "
		cadena += "\""
		cadena += raiz.B_info.Value.Date + "\n" + raiz.B_info.Value.State + "\n" + raiz.B_info.Value.Name + "\n" + strconv.Itoa(raiz.B_info.Value.Tutor)
		cadena += "\" [dir=back];\n "
	}
	return cadena
}
