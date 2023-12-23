# Manual Tecníco  :cat:

El proyecto consiste en la creacion de una aplicacion la cual sirve del lado del administrador para manejar informacion de estudiantes que quieran solicitar tutorias asi como para poder administrar solicitudes de tutores y a los estudiantes usuarios les permite asignarse a un curso mediante su codigo.

En el desarrollo del proyecto se hicieron uso de diversas herramientas:


  | |  |
  |---|---|
  | GO | Lenguaje que se utilizo para la creacion del proyecto |
  | Grapviz| es un software de codigo abierto el cual sirve  para visualizacion de graficas.|
  |vscode|Es un editor de codigo el cual es desarrollado por Microsoft. |
  |Git|es una forja para alojar proyectos utilizando el sistema de control de versiones Git. |

## Funcionalidades Principales

- Administrador
  * Archivo Estudiantes [.csv]: En esta opcion se permite al usuario ingresar el nombre del archivo donde se encuentra contenida la informacion de los estudiantes, la cual es guardada en una lista doble.
  
  ![listaDoble](data:image/jpeg;base64,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) 
  * Archivo Tutores [.csv]: Aqui se Ingresan los datos de los tutores, como son el nombre, carnet, curso que desea dar y su promedio. el cual es importante ya que dependiendo de el se le asignara un nivel de prioridad para su aprovacion de tutoria. Esta informacion se maneja en 2 fases por lo que primero se guarda la informacion en un cola de prioridad donde va de 1 a 4 y se ingresa a una pila.
  
    |Rango de nota|Prioridad|
    |-------------|---------|
    |   90 - 100  |   1     |
    |   75 - 89   |   2     |
    |   65 - 74   |   3     |
    |   61 - 64   |   4     |

  * Archivo cursos [.Json]: Se emplea esta opcion la cual permite ingresar un listado de cursos los cuales tienen los datos de nombre y codigo y se van guardando en un ArbolAVL.
  
  ![arbolAvl](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ9PEPUmF14PwYNn-VP9asP7vWyzTOZTlICsA&usqp=CAU)
  * Administrar solicitudes de tutores: aqui se permite al administrador aceptar o rechazar los tutores, aqui se muestran primero los tutores segun su prioridad, los datos se almacenan en una lista doble circular.


  * Reportes: se hace uso de la herramienta de Graphviz para graficar los datos mediante las estructuras que se utilzaron para almacenarlos, como pueden ser listas dobdes, Matrices dispersas, Arboles Avl.
  
  ![Grapviz](https://spin.atomicobject.com/wp-content/uploads/20170927122530/vscode-graphviz-live-preview.gif)



## Estructura de datos utilizadas:

se hizo uso de distintas estructuras las cuales sirvieron para almacenar los datos y asi poder analizarlos y almacenarlos:

* Listas Dobles: Son estructura que se caracterizan por poseer 2 punteros Anterior y Siguiente los caules apuntan al dato respectivo. Ventaja de usar listas es que son estructuras dinamicas lo que quiere decir que no almacenan una cantidad fija de datos, lo cual es importante ya que crece su espacio segun lo que se va consumiento.
![listaD2](https://ccia.ugr.es/~jfv/ed1/tedi/cdrom/icons/lenlaz1.gif)

``` Go
\\Creacion de un nodo y una lisat doble en Go

type NodoLD struct {
	Dato    String
	Next   *NodoLD
	Prev   *NodoLD
}


type ListD struct {
	First *NodoLD
	Len   int
}
```
* Cola: Se basan en el Principio Fifo (First in, first out) o Ultimo en entrar, primero en salir. 
  
![pila](https://upload.wikimedia.org/wikipedia/commons/thumb/b/bb/Cola.svg/1200px-Cola.svg.png)

```Golang
type Cola struct {
	First *NodoCola
	Len   int
}
```

* Matrices Dispersas
*  Arbol Avl

