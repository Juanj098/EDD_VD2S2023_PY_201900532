package main

import (
	"EDD_VD2S2023_PY_201900532/structs"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

var listaD *structs.ListD = &structs.ListD{First: nil, Len: 0}
var listaDC *structs.ListDC = &structs.ListDC{First: nil, Len: 0}

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

func main() {

	for {
		var Input int = 0
		fmt.Println("<---------Login-------->")
		fmt.Println("| 1. Inicio de sesion  |")
		fmt.Println("| 2. Salir             |")
		fmt.Println("<---------------------->")
		fmt.Print("-> Opcion: ")
		fmt.Scan(&Input)
		fmt.Println("<---------------------->")
		if Input != 0 {
			switch Input {
			case 1:
				var userAdmin string = "ADMIN_201900532"
				var user string
				var passAdmin string = "Admin"
				var pass string
				fmt.Print("Usuario: ")
				fmt.Scan(&user)
				fmt.Print("Password: ")
				fmt.Scan(&pass)
				if user == userAdmin && pass == passAdmin {
					MAdmin()
				} else {
					clearTerminal()
					fmt.Println("usuario no existe o credenciales incorrectas...")
				}
			case 2:
				clearTerminal()
				fmt.Println("saliendo...")
				os.Exit(0)

			default:
				clearTerminal()
				fmt.Println("opcion no disponible..")
			}
		} else {
			fmt.Println("Ingrese una opcion")
		}

	}

}

func MAdmin() {
	var exit bool = false
	clearTerminal()
	for !exit {
		var opc string
		fmt.Println("<------------[ADMIN]------------>")
		fmt.Println("|1. Cargar estudiantes Tutores  |")
		fmt.Println("|2. Cargar estudiantes          |")
		fmt.Println("|3. Cargar cursos al sistema    |")
		fmt.Println("|4. Control de Tutores          |")
		fmt.Println("|5. Reportes Estructuras        |")
		fmt.Println("|6. Cerrar Sesion               |")
		fmt.Println("<------------------------------->")
		fmt.Print("-> ")
		fmt.Scan(&opc)
		opcI, _ := strconv.Atoi(opc)
		fmt.Println("<------------------------------->")
		clearTerminal()
		switch opcI {
		case 1:
			Tutores()
		case 2:
			Estudiantes()
		case 3:
			fmt.Println("Cursos .Json")
		case 4:
			fmt.Println("Control Tutores .csv")
		case 5:
			fmt.Println("Reportes")
		case 6:
			exit = true
		default:
			fmt.Println("Opcion no disponible...")
		}
	}
}

func Estudiantes() {
	var exit bool = false
	clearTerminal()
	for !exit {
		var opc string
		fmt.Println("<------------------->")
		fmt.Println("|1. cargar .csv     |")
		fmt.Println("|2. ver estudiantes |")
		fmt.Println("|3. salir           |")
		fmt.Println("<------------------->")
		fmt.Print("-> ")
		fmt.Scan(&opc)
		opcI, _ := strconv.Atoi(opc)
		fmt.Println("<------------------->")
		clearTerminal()
		switch opcI {
		case 1:
			fmt.Println("leer Doc")
			var path string
			fmt.Print("Ingrese Nombre Documento: ")
			fmt.Scan(&path)
			listaD.ReadCSV(path)
		case 2:
			fmt.Println("-listAlumnos-")
			listaD.ViewList()
		case 3:
			exit = true
		default:
			fmt.Println("Opcion no disponible...")

		}
	}
}

func Tutores() {
	var exit bool = false
	clearTerminal()
	for !exit {
		var opc string
		fmt.Println("<------------------>")
		fmt.Println("|1. cargar .csv    |")
		fmt.Println("|2. ver tutores    |")
		fmt.Println("|3. salir          |")
		fmt.Println("<------------------>")
		fmt.Print("-> ")
		fmt.Scan(&opc)
		opcI, _ := strconv.Atoi(opc)
		fmt.Println("<------------------>")
		clearTerminal()
		switch opcI {
		case 1:
			fmt.Println("leer .csv")
			var path string
			fmt.Print("Ingrese Nombre Documento: ")
			fmt.Scan(&path)
			listaDC.ReadCSV(path)
		case 2:
			fmt.Println("-listTutores-")
			listaDC.ViewListC()
		case 3:
			exit = true
		default:
			fmt.Println("Opcion invalida...")
		}
	}
}
