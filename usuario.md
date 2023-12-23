# Manual de Usuario 



## Descricion del proyecto:

 La aplicacion consiste en un sofware el cual permita administrar datos que son proporcionados por el usuario como tambien da la oportunidad al estudiante de asignarse con un tutor en especifico.

## Funciones Principales de la aplicacion:

### Archivos de Entrada: El usuario puede proporcionnar 2 tipos de archivos con los datos necesarios. Los culales son archivos con extension ,csv y tipo.json
```Json

\\Json

{
    "cursos":[
        {
            "codigo":"736",
            "Nombre":"Analisis Probabilistico"
        },
        {
            "codigo":"772",
            "Nombre":"Estructura de datos"
        },
        {
            "codigo":"796",
            "Nombre":"Lenguajes formales y de programacion"
        },
        {
            "codigo":"773",
            "Nombre":"Manejo e Implementacion de archivos"
        },
        {
            "codigo":"722",
            "Nombre":"Teoria de sistemas"
        },
        {
            "codigo":"14",
            "Nombre":"Economia"
        },
        {
            "codigo":"281",
            "Nombre":"Sistemas operativos"
        },
        {
            "Codigo": "781",
            "Nombre": "Organizacion de Lenguajes De Compiladores 2"
        }
    ]
}
```

```csv
202103105,Luis Daniel Castellanos Betancourt,772,80
202103206,Kewin Maslovy Patzan Tzun,781,86
202103252,Kevin Estuardo Del Cid Quezada,775,90
202106003,Jhonatan Alexander Aguilar Reyes,771,85
202109750,Luis Antonio Castro Padilla,770,65
202110180,Juan Carlos Gonzalez Valdez,772,69

```
* Inicio de Sesion:
La aplicacion provee 2 perfiles para el usurio los cuales son administrador y estudiante.
![log](https://github.com/Juanj098/EDD_VD2S2023_PY_201900532/blob/main/img/Captura%20de%20pantalla%202023-12-23%20024552.png?raw=true)
### Administrador:
tiene permitido todo e manejo administrativo, donde las principales funciones son:

![admin](https://github.com/Juanj098/EDD_VD2S2023_PY_201900532/blob/main/img/Captura%20de%20pantalla%202023-12-23%20025253.png?raw=true)
* Ingresar Documento de Tutores: Ingresar Documento csv con informacion de tutores.
* Ingrear Documento de Estudiantes: Ingresar documentos csv con informacion de estudiantes que desean tutorias
* Cargar cursos al sistema: cargar archivo Json con lisado de cursos que estan disponibles para tutorias
* Control de tutores: Se aceptan o rechazan las solicitudes para poder impartir tutorias
  
  ![acceptTutor](https://github.com/Juanj098/EDD_VD2S2023_PY_201900532/blob/main/img/Captura%20de%20pantalla%202023-12-23%20030555.png?raw=true)

* Reportes: Se da 5 opciones de reportes las cuales cada una contiene la informacion antes proporcionada por el administrador o el usuario

    ![report](https://github.com/Juanj098/EDD_VD2S2023_PY_201900532/blob/main/img/Captura%20de%20pantalla%202023-12-23%20025342.png?raw=true)
### Usuario
 Permite al estudiante poder ver informacion de ciertos cursos y poder asignarse a tutorias sobre ellos

 ![student](https://github.com/Juanj098/EDD_VD2S2023_PY_201900532/blob/main/img/Captura%20de%20pantalla%202023-12-23%20030814.png?raw=true)

 * Infor cursos: Muestra al estudiante informacion de los cursos disponibles y muestra al tutor que lo imparte
  
    ![infoTuto](https://github.com/Juanj098/EDD_VD2S2023_PY_201900532/blob/main/img/Captura%20de%20pantalla%202023-12-23%20030845.png?raw=true)

* Asignar cursos: el estudiante ingresa el codigo del curso que desea asignarse y listo si esta disponible aparecera un mensaje de exito.
  