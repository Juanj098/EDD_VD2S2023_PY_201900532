package main

import (
	"I_Fase/structs"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

var listaD *structs.ListD = &structs.ListD{First: nil, Len: 0}
var listaDC *structs.ListDC = &structs.ListDC{First: nil, Len: 0}
var ColaPri *structs.Cola = &structs.Cola{First: nil, Len: 0}
var matriz *structs.Matriz = &structs.Matriz{Raiz: &structs.NodoMtx{PosX: -1, PosY: -1, Dato: &structs.Data{Ctutor: 0, Cestudiante: 0, Curso: "RAIZ"}}, CantAl: 0, CantTut: 0}
var Arbol *structs.ArbolAVL = &structs.ArbolAVL{Raiz: nil}

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
				passInt, _ := strconv.Atoi(pass)
				userLoad := listaD.Search(user, passInt)
				if user == userAdmin && pass == passAdmin {
					MAdmin()
				} else if userLoad != nil {
					clearTerminal()
					MStudent(userLoad)
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
			Cursos()
		case 4:
			CTtutores()
		case 5:
			Reportes()
		case 6:
			exit = true
		default:
			fmt.Println("Opcion no disponible...")
		}
	}
}

func MStudent(user *structs.Alumnos) {
	var exit bool = false
	clearTerminal()
	for !exit {
		var opc string
		fmt.Println("<------------[" + strconv.Itoa(user.Carnet) + "]------------>")
		fmt.Println("|1. Info cursos                     |")
		fmt.Println("|2. Asignar cursos                  |")
		fmt.Println("|3. Cerrar Sesion                   |")
		fmt.Println("<----------------------------------->")
		fmt.Print("-> ")
		fmt.Scan(&opc)
		opcI, _ := strconv.Atoi(opc)
		fmt.Println("<----------------------------------->")
		clearTerminal()
		switch opcI {
		case 1:
			fmt.Println("Informacion de Cursos")
			listaDC.ViewListC()
		case 2:
			fmt.Println("Asignar")
			var code_curso string
			fmt.Print("Ingrese codigo de curso que desea asignar: ")
			fmt.Scan(&code_curso)
			Asignar(code_curso, user.Carnet)
		case 3:
			exit = true
		default:
			fmt.Println("Opcion no disponible...")
		}
	}
}

func Asignar(curso string, log int) {
	if Arbol.Search(curso) {
		tutorS := listaDC.SearchTutor(curso)
		if tutorS != nil {
			// student, err := strconv.Atoi(log)
			// if err != nil {
			// 	return
			// }
			matriz.AddElm(log, tutorS.Carnet, curso)
			fmt.Println("se asigno de forma exitosa el curso!...")
		} else {
			fmt.Println("Curso no cuenta con tutor...")
		}
	} else {
		fmt.Println("Curso no disponible")
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
			fmt.Println("leer .csv -estudiantes-")
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
		fmt.Println("<-----{TUTORES}----->")
		fmt.Println("|1. cargar .csv     |")
		fmt.Println("|2. salir           |")
		fmt.Println("<------------------->")
		fmt.Print("-> ")
		fmt.Scan(&opc)
		opcI, _ := strconv.Atoi(opc)
		fmt.Println("<------------------->")
		clearTerminal()
		switch opcI {
		case 1:
			fmt.Println("leer .csv -tutores-")
			var path string
			fmt.Print("Ingrese Nombre Documento: ")
			fmt.Scan(&path)
			ColaPri.ReadCSV(path)
		case 2:
			exit = true
		default:
			fmt.Println("Opcion invalida...")
		}
	}
}

func CTtutores() {
	var exit bool = false
	clearTerminal()
	for !exit {
		var opc string
		fmt.Println("<---------{TUTORES}--------->")
		fmt.Println("|1. Administrar solicitudes |")
		fmt.Println("|2. aceptados               |")
		fmt.Println("|3. salir                   |")
		fmt.Println("<--------------------------->")
		fmt.Print("-> ")
		fmt.Scan(&opc)
		opcI, _ := strconv.Atoi(opc)
		fmt.Println("<--------------------------->")
		clearTerminal()
		switch opcI {
		case 1:
			var exit bool = false
			clearTerminal()
			for !exit {
				ColaPri.Primero()
				fmt.Println("-----------------------------------")
				fmt.Println("| (A) Aceptar                     |")
				fmt.Println("| (X) Rechazar                    |")
				fmt.Println("| (S) Salir                       |")
				fmt.Println("<--------------------------------->")
				fmt.Print("-> ")
				fmt.Scan(&opc)
				opc = strings.ToLower(opc)
				clearTerminal()
				switch opc {
				case "a":
					tutor := ColaPri.Ret_Desc()
					resp := listaDC.SearchTutor(tutor.Curso)
					if resp != nil {
						if tutor.Prom >= resp.Prom {
							listaDC.AddTutor(tutor.Name, tutor.Carnet, tutor.Curso, tutor.Prom)
							matriz.AddElm(tutor.Carnet, 0, tutor.Curso)
							fmt.Println("tutor registrado..")
						} else {
							fmt.Println("curso ya cuenta con tutor...")
						}
					} else {
						listaDC.AddTutor(tutor.Name, tutor.Carnet, tutor.Curso, tutor.Prom)
						matriz.AddElm(tutor.Carnet, 0, tutor.Curso)
						fmt.Println("tutor registrado..")
					}
				case "x":
					ColaPri.Descolar()
				case "s":
					exit = true
				default:
					fmt.Println("opcion no disponible...")
				}

			}
		case 2:
			listaDC.ViewListC()
		case 3:
			exit = true
		default:
			fmt.Println("Opcion invalida...")
		}
	}
}

func Cursos() {
	var exit bool = false
	clearTerminal()
	for !exit {
		var opc string
		fmt.Println("<----------{CURSOS}---------->")
		fmt.Println("|1. Registrar cursos         |")
		fmt.Println("|2. salir                    |")
		fmt.Println("<---------------------------->")
		fmt.Print("-> ")
		fmt.Scan(&opc)
		opcI, _ := strconv.Atoi(opc)
		clearTerminal()
		fmt.Println("<---------------------------->")
		switch opcI {
		case 1:
			fmt.Println("Cursos .Json")
			var path string
			fmt.Print("Ingrese nombre de Archivo JSON: ")
			fmt.Scan(&path)
			Arbol.ReadJSON(path)
		case 2:
			exit = true
		default:
			fmt.Println("Opcion no disponible...")
		}
	}
}

func Reportes() {
	var exit bool = false
	clearTerminal()
	for !exit {
		var opc string
		fmt.Println("<--------{REPORTES}--------->")
		fmt.Println("|1. Reporte Alumnos         |")
		fmt.Println("|2. Reporte Tutores         |")
		fmt.Println("|3. Matriz                  |")
		fmt.Println("|4. Arbol Cursos            |")
		fmt.Println("|5. salir                   |")
		fmt.Println("<--------------------------->")
		fmt.Print("-> ")
		fmt.Scan(&opc)
		opcI, _ := strconv.Atoi(opc)
		fmt.Println("<--------------------------->")
		clearTerminal()
		switch opcI {
		case 1:
			listaD.ReporteG()
		case 2:
			listaDC.ReporteG()
		case 3:
			matriz.ReporteG()
		case 4:
			Arbol.ReporteG()
		case 5:
			exit = true
		default:
			fmt.Println("Opcion no disponible...")
		}
	}
}
