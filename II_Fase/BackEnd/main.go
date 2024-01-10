package main

import (
	// "Backend/Structs"
	"BackEnd/Structs/AMerkle"
	grafos "BackEnd/Structs/Grafos"
	"BackEnd/Structs/THash"
	"BackEnd/Structs/arbolB"
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// structs
var ArbolB *arbolB.ArbolB = &arbolB.ArbolB{Raiz: nil, Orden: 3}
var lista *arbolB.ListaSimple = &arbolB.ListaSimple{Inicio: nil, Longitud: 0}
var Thash *THash.TablaHash = &THash.TablaHash{Tabla: make(map[int]THash.NodoHash), Capacidad: 7, Utilizacion: 0}
var ArbolM *AMerkle.ArbolMerkle = &AMerkle.ArbolMerkle{Root: nil, BData: nil, Cant: 0}
var grafo *grafos.Grafo = &grafos.Grafo{Principal: nil}

type User struct {
	UserName string `json:"userlog"`
	Password string `json:"Pass"`
}
type Reporte struct {
	Name       string `json:"Nombre"`
	Estructura string `json:"Estrucura_solicitada"`
}
type Response struct {
	Response string `json:"response"`
	User     string
}

type Cursos struct {
	Codigo string   `json:"codigo"`
	Post   []string `json:"Post"`
}

type dataStudent struct {
	Response string `json:"response"`
	User     string
	Data     []int
}

type Report struct {
	Estructura string
}
type CursosList struct {
	Cursos []Cursos `json:"Cursos"`
}

type BookState struct {
	Name   string
	Estado string
	Tutor  int
}

type CursoBooks struct {
	Curso int
}

type CarnetE struct {
	Carnet int
}

var userAdmin = User{
	UserName: "ADMIN_201900532",
	Password: "Admin",
}

func ValidateUser(c *fiber.Ctx) error {
	var p User
	var cursos []int
	if err := c.BodyParser(&p); err != nil {
		return err
	}
	if p.UserName == userAdmin.UserName {
		r := Response{Response: "Admin", User: userAdmin.UserName}
		return c.Status(fiber.StatusOK).JSON(r)
	} else {
		userlog := p.UserName
		if resp := ArbolB.Buscar(userlog, lista); resp != nil {
			if resp.Dato.Password == SHA256(p.Password) {
				response := Response{Response: "Tutor", User: strconv.Itoa(resp.Dato.Carnet)}
				return c.Status(fiber.StatusOK).JSON(response)
			} else {
				response := Response{Response: "Credenciales Invalidas..."}
				return c.Status(fiber.StatusOK).JSON(response)
			}
		}
		if resp := Thash.Buscar(userlog, SHA256(p.Password)); resp != nil {
			for i := 0; i < len(resp.Dato.Cursos); i++ {
				cursos = append(cursos, resp.Dato.Cursos[i])
			}
			response := dataStudent{Response: "Estudiante", User: strconv.Itoa(resp.Dato.Carnet), Data: cursos}
			return c.Status(fiber.StatusOK).JSON(response)
		} else {
			response := Response{Response: "Credenciales Invalidas..."}
			return c.Status(fiber.StatusOK).JSON(response)
		}
	}
}

func SHA256(cadena string) string {
	hexa := ""
	h := sha256.New()
	h.Write([]byte(cadena))
	hexa = hex.EncodeToString(h.Sum(nil))
	return hexa
}

func ReaderCsvT(filename string) error {
	csvfile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer csvfile.Close()
	r := csv.NewReader(csvfile)
	_, err = r.Read()
	if err != nil {
		log.Fatal(err)
	}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		carnet, _ := strconv.Atoi(record[0])
		nTutor := record[1]
		curso, _ := strconv.Atoi(record[2])
		pass := SHA256(record[3])
		ArbolB.InsertA(carnet, nTutor, curso, pass)
	}
	return nil
}

func searchCsvT(filename string) error {
	var pathFile string = "C:\\Users\\juanj\\OneDrive\\Escritorio\\EDD\\EDD_VD2S2023_PY_201900532\\II_Fase\\backend\\uploads"
	var path string
	err := filepath.Walk(pathFile, func(ruta string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".csv" && info.Name() == filename {
			path = ruta
			return nil
		}
		return nil
	})
	if err != nil {
		return err
	}
	if path == "" {
		return fmt.Errorf("no se encontro doc...")
	}
	err = ReaderCsvT(path)
	if err != nil {
		return err
	}
	return nil
}

func searchCsvE(filename string) error {
	var pathFile string = "C:\\Users\\juanj\\OneDrive\\Escritorio\\EDD\\EDD_VD2S2023_PY_201900532\\II_Fase\\backend\\uploads"
	var path string
	err := filepath.Walk(pathFile, func(ruta string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".csv" && info.Name() == filename {
			path = ruta
			return nil
		}
		return nil
	})
	if err != nil {
		return err
	}
	if path == "" {
		return fmt.Errorf("no se encontro doc...")
	}
	err = ReaderCsvE(path)
	if err != nil {
		return err
	}
	return nil
}

func ReaderCsvE(filename string) error {
	csvfile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer csvfile.Close()
	r := csv.NewReader(csvfile)
	_, err = r.Read()
	if err != nil {
		return err
	}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		fmt.Println(record)
		carnet, err := strconv.Atoi(record[0])
		name := record[1]
		pass := SHA256(record[2])
		c1 := record[3]
		c2 := record[4]
		c3 := record[5]
		if c1 != "" && c2 != "" && c3 != "" {
			c1, _ := strconv.Atoi(record[3])
			c2, _ := strconv.Atoi(record[4])
			c3, _ := strconv.Atoi(record[5])
			Thash.Insertar(carnet, name, pass, c1, c2, c3)
		}

	}
	return nil
}

func Tcsv(c *fiber.Ctx) error {
	file, err := c.FormFile("upload")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "error al obtener el archivo",
		})
	}
	errSave := c.SaveFile(file, "uploads/"+file.Filename)
	if errSave != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "error al guardar el archivo",
		})
	}
	searchCsvT(file.Filename)
	return c.JSON(fiber.Map{
		"success":  true,
		"message":  "archivo procesado con exito :D",
		"filename": file.Filename,
	})

}

func Ecsv(c *fiber.Ctx) error {
	file, err := c.FormFile("upload")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "error al obtener el archivo",
		})
	}
	errSave := c.SaveFile(file, "uploads/"+file.Filename)
	if errSave != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "error al guardar el archivo",
		})
	}
	searchCsvE(file.Filename)
	return c.JSON(fiber.Map{
		"success":  true,
		"message":  "archivo procesado con exito :D",
		"filename": file.Filename,
	})
}

func searchAB(c *fiber.Ctx) error {
	var p User
	err := c.BodyParser(&p)
	if err != nil {
		return err
	}
	userlog := p.UserName
	fmt.Println(userlog)
	resp := ArbolB.Buscar(userlog, lista)
	fmt.Println(resp)
	if resp != nil {
		// temp := lista.Buscar(userlog)
		return c.JSON(fiber.Map{
			"success": true,
			"message": resp.Dato.Name,
		})
	} else {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "no",
		})
	}
}

func searchTH(c *fiber.Ctx) error {
	var p User
	err := c.BodyParser(&p)
	if err != nil {
		return err
	}
	userlog := p.UserName
	fmt.Print(userlog)
	resp := Thash.Buscar(userlog, SHA256(p.Password))
	if resp != nil {
		return c.JSON(fiber.Map{
			"success": true,
			"message": resp.Dato.Name,
		})
	} else {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "no",
		})
	}
}

func Students(c *fiber.Ctx) error {
	arreglo := Thash.Arreglo()
	return c.JSON(arreglo)
}

func cJSON(c *fiber.Ctx) error {
	file, err := c.FormFile("upload")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "error al obtener el archivo",
		})
	}
	errSave := c.SaveFile(file, "uploads/"+file.Filename)
	if errSave != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "error al guardar el archivo",
		})
	}
	searchJSON(file.Filename)
	return c.JSON(fiber.Map{
		"success":  true,
		"message":  "archivo procesado con exito :D",
		"filename": file.Filename,
	})
}

func searchJSON(filename string) error {
	var pathFile string = "C:\\Users\\juanj\\OneDrive\\Escritorio\\EDD\\EDD_VD2S2023_PY_201900532\\II_Fase\\backend\\uploads"
	var path string
	err := filepath.Walk(pathFile, func(ruta string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".json" && info.Name() == filename {
			path = ruta
			return nil
		}
		return nil
	})
	if err != nil {
		return err
	}
	if path == "" {
		return nil
	}
	ReaderJson(path)
	return nil
}

func ReaderJson(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var cursosList CursosList
	json.Unmarshal(data, &cursosList)
	cursos := cursosList.Cursos
	for i := 0; i < len(cursos); i++ {
		for j := 0; j < len(cursos[i].Post); j++ {
			grafo.Insertar(cursos[i].Codigo, cursos[i].Post[j])
		}
	}
	fmt.Println("s")
}

func Tlibro(c *fiber.Ctx) error {
	var libros arbolB.Libros
	if err := c.BodyParser(&libros); err != nil {
		return c.JSON(fiber.Map{
			"error":   "true",
			"message": err,
		})
	}
	fmt.Println(libros)
	if ArbolB.Raiz == nil {
		return c.JSON(fiber.Map{
			"message": "Ingrese Tutores...",
		})
	}
	ArbolB.SaveBook(ArbolB.Raiz.First, libros.Name, libros.Tutor, libros.Estado)
	return c.JSON(fiber.Map{
		"message": "libro procesado con exito :D",
	})
}

func Tpublic(c *fiber.Ctx) error {
	var publicacion arbolB.Publicacion
	if err := c.BodyParser(&publicacion); err != nil {
		return c.JSON(fiber.Map{
			"error":   "true",
			"message": err,
		})
	}
	if ArbolB.Raiz == nil {
		return c.JSON(fiber.Map{
			"message": "Ingrese Tutores...",
		})
	}
	ArbolB.NewPublic(ArbolB.Raiz.First, publicacion.Tutor, publicacion.Publicacion)
	fmt.Println(publicacion)
	return c.JSON(fiber.Map{
		"message": "Publicacion registrada...",
	})
}

func AceptBooks(c *fiber.Ctx) error {
	LTemp := &arbolB.ListaSimple{Inicio: nil, Longitud: 0}
	var Books []*arbolB.Libros
	if ArbolB.Raiz != nil {
		ArbolB.RetBookPnedientes(ArbolB.Raiz.First, LTemp)
		if LTemp.Longitud > 0 {
			aux := LTemp.Inicio
			for aux != nil {
				for i := 0; i < len(aux.Tutor.Dato.Libros); i++ {
					if aux.Tutor.Dato.Libros[i].Estado == "Pendiente" {
						Books = append(Books, *&aux.Tutor.Dato.Libros[i])
					}
				}
				aux = aux.Siguiente
			}
		}

	}
	if len(Books) > 0 {
		return c.JSON(&fiber.Map{
			"status":           "200",
			"LibrosPendientes": Books,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 500,
	})
}

func estadoLibro(c *fiber.Ctx) error {
	var estado BookState
	if err := c.BodyParser(&estado); err != nil {
		return c.JSON(fiber.Map{
			"error":   "true",
			"message": err,
		})
	}
	fmt.Println(estado)
	ArbolM.AgregarB(estado.Estado, estado.Name, estado.Tutor)
	ArbolB.CambiarEstado(ArbolB.Raiz.First, estado.Name, estado.Estado, estado.Tutor)
	return c.JSON(fiber.Map{
		"message": "Estado modificado...",
	})
}

func agregarAM(c *fiber.Ctx) error {
	ArbolM.GenArbol()
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Arbol creado...",
	})
}

func searchBooks(c *fiber.Ctx) error {
	listS := &arbolB.ListaSimple{Inicio: nil, Longitud: 0}
	var C_libro CursoBooks
	var Books []arbolB.Libros
	if err := c.BodyParser(&C_libro); err != nil {
		return err
	}
	if ArbolB.Raiz != nil {
		ArbolB.SearchBooks(ArbolB.Raiz.First, C_libro.Curso, listS)
		if listS.Longitud > 0 {
			aux := listS.Inicio
			for aux != nil {
				for i := 0; i < len(aux.Tutor.Dato.Libros); i++ {
					if aux.Tutor.Dato.Libros[i].Estado == "Aceptado" {
						Books = append(Books, *aux.Tutor.Dato.Libros[i])
					}
				}
				aux = aux.Siguiente
			}
		}
	}
	if len(Books) > 0 {
		return c.JSON(&fiber.Map{
			"status":          "200",
			"LibrosAceptados": Books,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 500,
	})
}

func searchPublicaciones(c *fiber.Ctx) error {
	listS := &arbolB.ListaSimple{Inicio: nil, Longitud: 0}
	var C_libro CursoBooks
	var Publicacion []arbolB.Publicacion
	if err := c.BodyParser(&C_libro); err != nil {
		return err
	}
	if ArbolB.Raiz != nil {
		ArbolB.SearchPublicaciones(ArbolB.Raiz.First, C_libro.Curso, listS)
		fmt.Println(listS.Longitud)
		if listS.Longitud > 0 {
			aux := listS.Inicio
			for aux != nil {
				for t := 0; t < len(aux.Tutor.Dato.Publicacion); t++ {
					if aux.Tutor.Dato.Publicacion[t].Publicacion != "" {
						Publicacion = append(Publicacion, *aux.Tutor.Dato.Publicacion[t])
					}
				}
				aux = aux.Siguiente
			}

		}
	}
	if len(Publicacion) > 0 {
		return c.JSON(&fiber.Map{
			"status":        "200",
			"Publicaciones": Publicacion,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 500,
	})
}

func ObetenerC(c *fiber.Ctx) error {
	var carnet CarnetE
	if err := c.BodyParser(&carnet); err != nil {
		return err
	}
	search := Thash.BuscarC(carnet.Carnet)
	if search != nil {
		return c.JSON(&fiber.Map{
			"status":  200,
			"Arreglo": search.Dato.Cursos,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 500,
	})
}

func ReporteG(c *fiber.Ctx) error {
	var estructura Report
	if err := c.BodyParser(&estructura); err != nil {
		return err
	}
	if estructura.Estructura == "ArbolB" {
		ArbolB.Graficar()
	} else if estructura.Estructura == "Tabla-Hash" {
		return nil
	} else if estructura.Estructura == "Grafo" {
		grafo.Reporte()
	} else if estructura.Estructura == "Arbol-Merkle" {
		ArbolM.Graficar()
	}
	return c.JSON(&fiber.Map{
		"status":  200,
		"Arreglo": "estructura creada",
	})
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("http://localhost:3000/Login")
	})

	admin := app.Group("/Admin")
	Tutor := app.Group("/Tutor")
	Estudiante := app.Group("/Estudiante")
	admin.Post("/TutoresCsv", Tcsv)
	admin.Post("/EstudiantesCsv", Ecsv)
	admin.Get("/GetStudents", Students)
	admin.Post("/CursosJSON", cJSON)
	admin.Post("/bookState", estadoLibro)
	admin.Get("/LibrosPendientes", AceptBooks)
	admin.Get("/EndBooks", agregarAM)
	admin.Post("/Reporte", ReporteG)
	Tutor.Post("/CargarLibro", Tlibro)
	Tutor.Post("/CargarPublicacion", Tpublic)
	Estudiante.Post("/Books", searchBooks)
	Estudiante.Post("/Cursos", ObetenerC)
	Estudiante.Post("/Publicaciones", searchPublicaciones)
	app.Post("/Login", ValidateUser)
	app.Get("/PruebaB", searchAB)
	app.Get("/PruebaTH", searchTH)

	app.Listen(":3000")
}
