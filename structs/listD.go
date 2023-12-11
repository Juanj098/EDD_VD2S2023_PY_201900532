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
