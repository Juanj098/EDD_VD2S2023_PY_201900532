package structs

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type Cola struct {
	First *NodoCola
	Len   int
}

func (c *Cola) Queue(name string, carne int, curso string, promedio int) {
	newTutor := &Tutor{Carnet: carne, Name: name, Curso: curso, Prom: promedio}
	newNodo := &NodoCola{Tutor: newTutor, Sig: nil, Prioridad: 0}
	/*
		prioridad -> 1: 90 >= prom <= 100
					 2: 75 >= prom <= 89
					 3: 65 >= prom <= 74
					 4: 61 >= prom <= 64
	*/
	if promedio >= 90 && promedio <= 100 {
		newNodo.Prioridad = 1
	} else if promedio >= 75 && promedio <= 89 {
		newNodo.Prioridad = 2
	} else if promedio >= 65 && promedio <= 74 {
		newNodo.Prioridad = 3
	} else if promedio >= 61 && promedio <= 64 {
		newNodo.Prioridad = 4
	} else {
		newNodo.Prioridad = 5
	}

	if c.Len == 0 {
		c.First = newNodo
		c.Len++
	} else {
		aux := c.First
		for aux.Sig != nil {
			if aux.Sig.Prioridad > newNodo.Prioridad && aux.Prioridad == newNodo.Prioridad {
				newNodo.Sig = aux.Sig
				aux.Sig = newNodo
				c.Len++
				return
			} else if aux.Sig.Prioridad > newNodo.Prioridad && aux.Prioridad < newNodo.Prioridad {
				newNodo.Sig = aux.Sig
				aux.Sig = newNodo
				c.Len++
				return
			} else {
				aux = aux.Sig

			}
		}
		aux.Sig = newNodo
		c.Len++
	}

}

func (c *Cola) Descolar() {
	if c.Len == 0 {
		fmt.Println("cola vacia")
	} else {
		c.First = c.First.Sig
		c.Len--
	}
}

func (c *Cola) Ret_Desc() *Tutor {
	if c.Len == 0 {
		fmt.Println("cola vacia")
		return nil
	} else {
		obj := c.First.Tutor
		c.First = c.First.Sig
		c.Len--
		return obj
	}
}

func (c *Cola) Primero() {
	if c.Len == 0 {
		fmt.Println("pila vacia")
	} else {
		fmt.Println("<---------------------------------->")
		fmt.Println("|Carnet:", c.First.Tutor.Carnet)
		fmt.Println("|Nombre:", c.First.Tutor.Name)
		fmt.Println("|Promedio:", c.First.Tutor.Prom)
		fmt.Println("|Curso:", c.First.Tutor.Curso)
		fmt.Println("<--------------------------------->")
		if c.First.Sig != nil {
			fmt.Println("|siguiente", c.First.Sig.Tutor.Carnet)
		} else {
			fmt.Println(" -> No hay mas tutores")
		}
		fmt.Println("-----------------------------------")

	}
}

func (c *Cola) ReadCSV(path string) {
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
		carneCSV, _ := strconv.Atoi(record[0])
		cursoCSV := record[2]
		promCSV, _ := strconv.Atoi(record[3])
		c.Queue(nameCSV, carneCSV, cursoCSV, promCSV)
		// fmt.Println(record)
	}
}

func clearTerminal() {
	var cmd *exec.Cmd
	// Verifica el sistema operativo y selecciona el comando adecuado
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		fmt.Println("No se pudo determinar el sistema operativo.")
		return
	}
	// Ejecuta el comando para limpiar la terminal
	cmd.Stdout = os.Stdout
	cmd.Run()
}
