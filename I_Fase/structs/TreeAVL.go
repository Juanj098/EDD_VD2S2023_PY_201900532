package structs

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type ArbolAVL struct {
	Raiz *NodoTree
}

type Curso struct {
	Name   string `json:"Nombre"`
	Codigo string `json:"Codigo"`
}

type DataCurso struct {
	Cursos []Curso `json:"cursos"`
}

func (a *ArbolAVL) Altura(raiz *NodoTree) int {
	if raiz == nil {
		return 0
	}
	return raiz.Altura
}

func (a *ArbolAVL) Equilibrio(raiz *NodoTree) int {
	if raiz == nil {
		return 0
	} else {
		return (a.Altura(raiz.Right) - a.Altura(raiz.Left))
	}
}

func (a *ArbolAVL) RotacionIzq(raiz *NodoTree) *NodoTree {
	raiz_D := raiz.Right
	child_Izq := raiz_D.Left
	raiz_D.Left = raiz
	raiz.Right = child_Izq
	Max := math.Max(float64(a.Altura(raiz.Left)), float64(a.Altura(raiz.Right)))
	raiz.Altura = 1 + int(Max)
	raiz.F_equilibrio = a.Equilibrio(raiz)
	Max = math.Max(float64(a.Altura(raiz_D.Left)), float64(a.Altura(raiz_D.Right)))
	raiz_D.Altura = 1 + int(Max)
	raiz_D.F_equilibrio = a.Equilibrio(raiz_D)
	return raiz_D
}

func (a *ArbolAVL) RotacionDer(raiz *NodoTree) *NodoTree {
	raiz_Izq := raiz.Left
	child_Der := raiz_Izq.Right
	raiz_Izq.Right = raiz
	raiz.Left = child_Der
	Max := math.Max(float64(a.Altura(raiz.Left)), float64(a.Altura(raiz.Right)))
	raiz.Altura = 1 + int(Max)
	raiz.F_equilibrio = a.Equilibrio(raiz)
	Max = math.Max(float64(a.Altura(raiz_Izq.Left)), float64(a.Altura(raiz_Izq.Right)))
	raiz_Izq.Altura = 1 + int(Max)
	raiz_Izq.F_equilibrio = a.Equilibrio(raiz_Izq)
	return raiz_Izq
}

func (a *ArbolAVL) InsertarNodo(raiz *NodoTree, newNodo *NodoTree) *NodoTree {
	if raiz == nil {
		raiz = newNodo
	} else {
		if raiz.Value > newNodo.Value {
			raiz.Left = a.InsertarNodo(raiz.Left, newNodo)
		} else {
			raiz.Right = a.InsertarNodo(raiz.Right, newNodo)
		}
	}
	Max := math.Max(float64(a.Altura(raiz.Left)), float64(a.Altura(raiz.Right)))
	raiz.Altura = 1 + int(Max)
	balanceo := a.Equilibrio(raiz)
	raiz.F_equilibrio = balanceo
	if balanceo > 1 && newNodo.Value > raiz.Right.Value {
		// rotacion simple Izq
		return a.RotacionIzq(raiz)
	} else if balanceo < -1 && newNodo.Value < raiz.Left.Value {
		// rotacion simple Der
		return a.RotacionDer(raiz)
	} else if balanceo > 1 && newNodo.Value < raiz.Right.Value {
		// rotacion Doble Izq
		raiz.Right = a.RotacionDer(raiz.Right)
		return a.RotacionIzq(raiz)
	} else if balanceo < -1 && newNodo.Value > raiz.Left.Value {
		// rotacion Doble Der
		raiz.Left = a.RotacionIzq(raiz.Left)
		return a.RotacionDer(raiz)
	}
	return raiz
}

func (a *ArbolAVL) InsertElm(valor string) {
	newNodo := &NodoTree{Value: valor}
	a.Raiz = a.InsertarNodo(a.Raiz, newNodo)
}

func (a *ArbolAVL) SearchArbol(curso string, raiz *NodoTree) *NodoTree {
	var ValueFind *NodoTree
	if raiz != nil {
		if raiz.Value == curso {
			ValueFind = raiz
		} else {
			if raiz.Value > curso {
				ValueFind = a.SearchArbol(curso, raiz.Left)
			} else {
				ValueFind = a.SearchArbol(curso, raiz.Right)
			}
		}
	}
	return ValueFind
}

func (a *ArbolAVL) Search(curso string) bool {
	SearchElm := a.SearchArbol(curso, a.Raiz)
	if SearchElm != nil {
		return true
	}
	return false
}

func (a *ArbolAVL) ReporteG() {
	dotContent := ""
	fileDot := "Arbol.dot"
	fileImg := "Arbol.jpg"
	if a.Raiz != nil {
		dotContent += "digraph arbol{ "
		dotContent += a.DotContent(a.Raiz, 0)
		dotContent += "}"
	} else {
		fmt.Println("Arbol sin cursos...")
	}
	crearArchivo(fileDot)
	escribirArchivo(dotContent, fileDot)
	ejecutar(fileImg, fileDot)
}

func (a *ArbolAVL) DotContent(raiz *NodoTree, indice int) string {
	cadena := ""
	numero := indice + 1
	if raiz != nil {
		cadena += "\""
		cadena += raiz.Value
		cadena += "\" ;"
		if raiz.Left != nil && raiz.Right != nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Value
			cadena += "\" -> "
			cadena += a.DotContent(raiz.Left, numero)
			cadena += "\""
			cadena += raiz.Value
			cadena += "\" -> "
			cadena += a.DotContent(raiz.Right, numero)
			cadena += "{rank=same" + "\"" + (raiz.Left.Value) + "\"" + " -> " + "\"" + (raiz.Right.Value) + "\"" + " [style=invis]}; "
		} else if raiz.Left != nil && raiz.Right == nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Value
			cadena += "\" -> "
			cadena += a.DotContent(raiz.Left, numero)
			cadena += "\""
			cadena += raiz.Value
			cadena += "\" -> "
			cadena += "x" + strconv.Itoa(numero) + "[style=invis]"
			cadena += "{rank=same" + "\"" + (raiz.Left.Value) + "\"" + " -> " + "x" + strconv.Itoa(numero) + " [style=invis]}; "
		} else if raiz.Left == nil && raiz.Right != nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += raiz.Value
			cadena += "\" -> "
			cadena += "x" + strconv.Itoa(numero) + "[style=invis]"
			cadena += "; \""
			cadena += raiz.Value
			cadena += "\" -> "
			cadena += a.DotContent(raiz.Right, numero)
			cadena += "{rank=same" + " x" + strconv.Itoa(numero) + " -> \"" + (raiz.Right.Value) + "\"" + " [style=invis]}; "
		}
	}
	return cadena
}

func (a *ArbolAVL) ReadJSON(path string) {
	js, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	var Datos DataCurso
	errJ := json.Unmarshal(js, &Datos)
	if errJ != nil {
		log.Fatal(err)
	}
	for _, t := range Datos.Cursos {
		a.InsertElm(t.Codigo)
		// fmt.Println(t.Name)
	}
}
