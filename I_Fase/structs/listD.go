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

type ListD struct {
	First *NodoLD
	Len   int
}

func (l *ListD) AddAlumno(carnet int, name string) {
	newAlumno := &Alumnos{Name: name, Carnet: carnet}
	newNodo := &NodoLD{Alumno: newAlumno, Next: nil, Prev: nil}
	if l.First == nil {
		l.First = newNodo
		l.Len++
	} else {
		aux := l.First
		for aux.Next != nil {
			aux = aux.Next
		}
		newNodo.Prev = aux
		aux.Next = newNodo
		l.Len++
	}
}

func (l *ListD) Search(user string, carnet int) *Alumnos {
	if l.First == nil {
		fmt.Println("lista vacia...")
		return nil
	} else {
		aux := l.First
		for aux.Next != nil {
			if strconv.Itoa(aux.Alumno.Carnet) == user && aux.Alumno.Carnet == carnet {
				return aux.Alumno
			}
			aux = aux.Next
		}
	}
	return nil
}

func (l *ListD) ViewList() {
	if l.First == nil {
		fmt.Println("Lista vacia")
	} else {
		aux := l.First

		for aux.Next != nil {
			fmt.Println("nombre", " -> ", aux.Alumno.Name, " | ", "carnet", " -> ", aux.Alumno.Carnet)
			aux = aux.Next
		}
		fmt.Println("nombre", " -> ", aux.Alumno.Name, " | ", "carnet", " -> ", aux.Alumno.Carnet)
	}
}

func (l *ListD) ReadCSV(path string) {
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
		nameCSV := record[1]
		record[0] = strings.TrimPrefix(record[0], "\ufeff")
		carnetCSV, errcar := strconv.Atoi(record[0])
		if errcar != nil {
			fmt.Println(errcar)
			continue
		}
		l.AddAlumno(carnetCSV, nameCSV)
	}
}

func (l *ListD) ReporteG() {
	fileName := "./LDoble.dot"
	nameImg := "./LDoble.jpg"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	dotContent := `
digraph G {
	rankdir = LR
	node[shape=box]
	nullFirst[shape ="point"];
	nullLast[shape = "point"];
`
	arrows := ``
	// nodos
	if l.First == nil {
		fmt.Println("Lista vacia")
	} else {
		// rev := l.Len
		cont := 1
		aux := l.First
		arrows += "\tnullFirst -> nodo1[arrowhead=none]\n"
		for aux.Next != nil {
			dotContent += "\tnodo" + strconv.Itoa(cont) + "[label=" + strconv.Itoa(aux.Alumno.Carnet) + "];\n"
			aux = aux.Next
			cont++
		}
		dotContent += "\tnodo" + strconv.Itoa(cont) + "[label=" + strconv.Itoa(aux.Alumno.Carnet) + "];\n"
		for i := 1; i < l.Len; i++ {
			c := i + 1
			arrows += "\tnodo" + strconv.Itoa(i) + " -> " + "nodo" + strconv.Itoa(c) + ";\n"
			arrows += "\tnodo" + strconv.Itoa(c) + " -> " + "nodo" + strconv.Itoa(i) + ";\n"
			cont = c
		}
		arrows += "\tnodo" + strconv.Itoa(l.Len) + "-> nullLast[arrowhead=none]\n"
		dotContent += arrows
	}
	dotContent += "}"
	crearArchivo(fileName)
	escribirArchivo(dotContent, fileName)
	ejecutar(nameImg, fileName)
}
