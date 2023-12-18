package structs

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type ListDC struct {
	First *NodoLDC
	Len   int
}

func (l *ListDC) AddTutor(name string, carnet int, curso string, prom int) {
	newTutor := &Tutor{Carnet: carnet, Name: name, Curso: curso, Prom: prom}
	newNodo := &NodoLDC{Tutor: newTutor, Next: nil, Prev: nil}
	if l.First == nil {
		l.First = newNodo
		newNodo.Next = newNodo
		newNodo.Prev = newNodo
		l.Len++
	} else {
		last := l.First.Prev
		last.Next = newNodo
		newNodo.Prev = last
		newNodo.Next = l.First
		l.First.Prev = newNodo
		l.Len++
	}
}

func (l *ListDC) AddTutorSort(name string, carnet int, curso string, prom int) {
	newTutor := &Tutor{Carnet: carnet, Name: name, Curso: curso, Prom: prom}
	newNodo := &NodoLDC{Tutor: newTutor, Next: nil, Prev: nil}
	if l.First == nil {
		l.First = newNodo
		l.First.Next = newNodo
		l.First.Prev = newNodo
		l.Len++
	} else {
		aux := l.First
		cont := 1
		for cont < l.Len {
			if l.First.Tutor.Carnet > carnet {
				newNodo.Next = l.First
				newNodo.Prev = l.First
				l.First.Prev = newNodo
				l.First = newNodo
				l.Len++
				return
			}
			if aux.Tutor.Carnet < carnet {
				aux = aux.Next
			} else {
				newNodo.Prev = aux.Prev
				aux.Prev.Next = newNodo
				newNodo.Next = aux
				aux.Prev = newNodo
				l.Len++
				return
			}
			cont++
		}
		if aux.Tutor.Carnet > carnet {
			newNodo.Next = aux
			newNodo.Prev = aux.Prev
			aux.Prev.Next = newNodo
			aux.Prev = newNodo
			l.Len++
			return
		}
		newNodo.Prev = aux
		newNodo.Next = l.First
		aux.Next = newNodo
		l.Len++
	}
}

func (l *ListDC) ViewListC() {
	if l.First == nil {
		fmt.Println("Lista Vacia")
	} else {
		aux := l.First
		for {
			fmt.Println("carnet: ", aux.Tutor.Carnet, " -> ", "Tutor: ", aux.Tutor.Name)
			aux = aux.Next
			if aux == l.First {
				break
			}
		}
	}
}

func (l *ListDC) ReadCSV(path string) {
	csvfile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(csvfile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(record[0])
		nameCSV := record[1]
		record[0] = strings.TrimPrefix(record[0], "\ufeff")
		carnetCSV, _ := strconv.Atoi(record[0])
		cursoCSV := record[2]
		promCSV, _ := strconv.Atoi(record[3])
		l.AddTutor(nameCSV, carnetCSV, cursoCSV, promCSV)
	}
}

func (l *ListDC) ReporteG() {
	fileName := "./LDobleC.dot"
	nameImg := "./LDobleC.jpg"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	dotContent := `
digraph G {
	rankdir = LR
	node[shape=box]

`
	arrows := ``
	// nodos
	if l.First == nil {
		fmt.Println("Lista vacia")
	} else {
		// rev := l.Len
		cont := 0
		aux := l.First
		arrows += "\tnodo0 ->" + "nodo" + strconv.Itoa(l.Len-1) + "\n"
		for i := 0; i < l.Len; i++ {
			dotContent += "\tnodo" + strconv.Itoa(i) + "[label=" + strconv.Itoa(aux.Tutor.Carnet) + "];\n"
			aux = aux.Next
		}
		dotContent += "\tnodo" + strconv.Itoa(cont) + "[label=" + strconv.Itoa(aux.Tutor.Carnet) + "];\n"
		for i := 0; i < l.Len-1; i++ {
			c := i + 1
			arrows += "\tnodo" + strconv.Itoa(i) + " -> " + "nodo" + strconv.Itoa(c) + ";\n"
			arrows += "\tnodo" + strconv.Itoa(c) + " -> " + "nodo" + strconv.Itoa(i) + ";\n"
			cont = c
		}
		arrows += "\tnodo" + strconv.Itoa(l.Len-1) + "-> nodo0\n"
		dotContent += arrows
	}
	dotContent += "}"
	crearArchivo(fileName)
	escribirArchivo(dotContent, fileName)
	ejecutar(nameImg, fileName)
}
