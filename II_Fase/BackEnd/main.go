package main

import (
	// "Backend/Structs"

	"BackEnd/Structs/THash"
	"BackEnd/Structs/arbolB"
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
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
}

var userAdmin = User{
	UserName: "ADMIN_201900532",
	Password: "Admin",
}

func ValidateUser(c *fiber.Ctx) error {
	var p User
	if err := c.BodyParser(&p); err != nil {
		return err
	}
	if p.UserName == userAdmin.UserName {
		r := Response{Response: "Admin"}
		return c.Status(fiber.StatusOK).JSON(r)
	} else {
		userlog := p.UserName
		if resp := ArbolB.Buscar(userlog, lista); resp != nil {
			if resp.Dato.Password == SHA256(p.Password) {
				response := Response{Response: "Tutor"}
				return c.Status(fiber.StatusOK).JSON(response)
			} else {
				response := Response{Response: "Credenciales Invalidas..."}
				return c.Status(fiber.StatusOK).JSON(response)
			}
		}
		if resp := Thash.Buscar(userlog, SHA256(p.Password)); resp != nil {
			response := Response{Response: "Estudiante"}
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
		curso := record[2]
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
		Thash.Insertar(carnet, name, pass)
	}
	fmt.Println(Thash.Capacidad)
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
func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("http://localhost:3000/Login")
	})

	admin := app.Group("/Admin")
	admin.Post("/TutoresCsv", Tcsv)
	admin.Post("/EstudiantesCsv", Ecsv)
	app.Post("/Login", ValidateUser)
	app.Get("/PruebaB", searchAB)
	app.Get("/PruebaTH", searchTH)
	app.Listen(":3000")
}
