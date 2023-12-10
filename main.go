package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

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
		fmt.Println("<---------Login-------->")
		fmt.Println("| 1. Inicio de sesion  |")
		fmt.Println("| 2. Salir             |")
		fmt.Println("<---------------------->")
		fmt.Print("-> Opcion: ")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		// scan.Text()
		fmt.Println("<---------------------->")
		choice, _ := strconv.Atoi(scan.Text())
		switch choice {
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
					switch opcI {
					case 1:
						fmt.Println("Tutores .csv")
					case 2:
						fmt.Println("Estudiantes .csv")
					case 3:
						fmt.Println("Cursos .Json")
					case 4:
						fmt.Println("Control Tutores .csv")
					case 5:
						fmt.Println("Reportes")
					case 6:
						fmt.Println("saliendo...")
						exit = true
					default:
						fmt.Println("Opcion no disponible...")
					}
				}
			} else {
				fmt.Println("incorrecto...")
			}
		case 2:
			fmt.Println("saliendo...")
			os.Exit(0)
		}
	}

}
