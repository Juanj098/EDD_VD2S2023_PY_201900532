package structs

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Matriz struct {
	Raiz    *NodoMtx
	CantAl  int
	CantTut int
}

func (m *Matriz) Bcolumna(Carne_tutor int, curso string) *NodoMtx {
	aux := m.Raiz
	for aux != nil {
		if aux.Dato.Ctutor == Carne_tutor && aux.Dato.Curso == curso {
			return aux
		}
		aux = aux.Next
	}
	return nil
}

func (m *Matriz) Bfila(Carne_Estudiante int) *NodoMtx {
	aux := m.Raiz
	for aux != nil {
		if aux.Dato.Cestudiante == Carne_Estudiante {
			return aux
		}
		aux = aux.Down
	}
	return nil
}

func (m *Matriz) AddCol(newNodo *NodoMtx, nodoRaiz *NodoMtx) *NodoMtx {
	temp := nodoRaiz
	piv := false
	for {
		if temp.PosX == newNodo.PosX {
			temp.PosY = newNodo.PosY
			temp.Dato = newNodo.Dato
			return temp
		} else if temp.PosX > newNodo.PosX {
			piv = true
			break
		}
		if temp.Next != nil {
			temp = temp.Next
		} else {
			break
		}
	}
	if piv {
		newNodo.Next = temp
		newNodo.Prev = temp.Prev
		temp.Prev.Next = newNodo
		temp.Prev = newNodo
	} else {
		newNodo.Prev = temp
		temp.Next = newNodo
	}
	return newNodo
}

func (m *Matriz) newColumna(x int, carnetT int, curso string) *NodoMtx {
	newNodo := &NodoMtx{PosX: x, PosY: -1, Dato: &Data{Ctutor: carnetT, Cestudiante: 0, Curso: curso}}
	col := m.AddCol(newNodo, m.Raiz)
	return col
}

func (m *Matriz) AddFil(newNodo *NodoMtx, nodoRaiz *NodoMtx) *NodoMtx {
	temp := nodoRaiz
	piv := false
	for {
		if temp.PosY == newNodo.PosY {
			temp.PosX = newNodo.PosX
			temp.Dato = newNodo.Dato
			return temp
		} else if temp.PosY > newNodo.PosY {
			piv = true
			break
		} else {
			break
		}
	}
	if piv {
		newNodo.Down = temp
		newNodo.Up = temp.Up
		temp.Up.Down = newNodo
		temp.Up = newNodo
	} else {
		newNodo.Up = temp
		temp.Down = newNodo
	}
	return newNodo
}

func (m *Matriz) newFila(y int, carnetE int, curso string) *NodoMtx {
	newNodo := &NodoMtx{PosX: -1, PosY: y, Dato: &Data{Ctutor: 0, Cestudiante: carnetE, Curso: curso}}
	Fil := m.AddFil(newNodo, m.Raiz)
	return Fil
}

func (m *Matriz) AddElm(carnetTut int, carnetEs int, curso string) {
	nodoCol := m.Bcolumna(carnetTut, curso)
	nodoFil := m.Bfila(carnetEs)
	if nodoCol == nil && nodoFil == nil {
		m.CantAl++
		m.CantTut++
		nodoCol = m.newColumna(m.CantTut, carnetTut, curso)
		nodoFil = m.newFila(m.CantAl, carnetEs, curso)
		newNodo := &NodoMtx{PosX: nodoCol.PosX, PosY: nodoFil.PosY, Dato: &Data{Ctutor: carnetTut, Cestudiante: carnetEs, Curso: curso}}
		newNodo = m.AddCol(newNodo, nodoFil)
		newNodo = m.AddFil(newNodo, nodoCol)
	} else if nodoCol != nil && nodoFil == nil {
		m.CantAl++
		nodoFil = m.newFila(m.CantAl, carnetEs, curso)
		newNodo := &NodoMtx{PosX: nodoCol.PosX, PosY: nodoFil.PosY, Dato: &Data{Ctutor: carnetTut, Cestudiante: carnetEs, Curso: curso}}
		newNodo = m.AddCol(newNodo, nodoFil)
		newNodo = m.AddFil(newNodo, nodoCol)
	} else if nodoCol == nil && nodoFil != nil {
		m.CantTut++
		nodoCol = m.newColumna(m.CantTut, carnetTut, curso)
		newNodo := &NodoMtx{PosX: nodoCol.PosX, PosY: nodoFil.PosY, Dato: &Data{Ctutor: carnetTut, Cestudiante: carnetEs, Curso: curso}}
		newNodo = m.AddCol(newNodo, nodoFil)
		newNodo = m.AddFil(newNodo, nodoCol)
	} else if nodoCol != nil && nodoFil != nil {
		newNodo := &NodoMtx{PosX: nodoCol.PosX, PosY: nodoFil.PosY, Dato: &Data{Ctutor: carnetTut, Cestudiante: carnetEs, Curso: curso}}
		newNodo = m.AddCol(newNodo, nodoFil)
		newNodo = m.AddFil(newNodo, nodoCol)
	} else {
		fmt.Println("Error")
	}
}

func (m *Matriz) ReporteG() {
	fileDot := "matriz.dot"
	fileImg := "matriz.jpg"
	aux1 := m.Raiz
	aux2 := m.Raiz
	aux3 := m.Raiz

	file, err := os.Create(fileDot)
	if err != nil {
		log.Fatal(err)

	}
	var dotContent string
	defer file.Close()
	if aux1 != nil {
		dotContent = `
digraph G {
	node[shape=cluster]
	rankdir=UD
	{rank = min
`
		for aux1 != nil {
			dotContent += "\t\tnodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + "[label=\"" + strconv.Itoa(aux1.Dato.Ctutor) + "\" rankdir=LR" + " group=\"" + strconv.Itoa(aux1.PosX+1) + "\"]\n"
			aux1 = aux1.Next
		}
		dotContent += "}\n"
		aux2 = aux2.Down
		for aux2 != nil {
			aux1 = aux2
			dotContent += "\t{rank=same\n"
			flag_raiz := true
			for aux1 != nil {
				if flag_raiz {
					dotContent += "\t\tnodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + "[label=\"" + strconv.Itoa(aux1.Dato.Cestudiante) + "\" group=\"" + strconv.Itoa(aux1.PosX+1) + "\"]\n"
					flag_raiz = false
				} else {
					dotContent += "\t\tnodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + "[label=\"" + aux1.Dato.Curso + "\" group=\"" + strconv.Itoa(aux1.PosX+1) + "\"]\n"
				}
				aux1 = aux1.Next
			}
			dotContent += "}\n"
			aux2 = aux2.Down
		}
		aux2 = aux3
		for aux2 != nil {
			aux1 = aux2
			for aux1.Next != nil {
				dotContent += "\tnodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + " -> " + "nodo" + strconv.Itoa(aux1.Next.PosX+1) + strconv.Itoa(aux1.Next.PosY+1) + "[arrowhead=none]\n"
				aux1 = aux1.Next
			}
			aux2 = aux2.Down
		}
		aux2 = aux3
		for aux2 != nil {
			aux1 = aux2
			for aux1.Down != nil {
				dotContent += "\tnodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + " -> " + "nodo" + strconv.Itoa(aux1.Down.PosX+1) + strconv.Itoa(aux1.Down.PosY+1) + "[arrowhead=none]\n"
				aux1 = aux1.Down
			}
			aux2 = aux2.Next
		}
		dotContent += "}"
	} else {
		fmt.Println("No hay elementos en la matriz...")
	}
	crearArchivo(fileDot)
	escribirArchivo(dotContent, fileDot)
	ejecutar(fileImg, fileDot)
}
