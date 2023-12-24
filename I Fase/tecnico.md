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
  
  
  * Archivo Tutores [.csv]: Aqui se Ingresan los datos de los tutores, como son el nombre, carnet, curso que desea dar y su promedio. el cual es importante ya que dependiendo de el se le asignara un nivel de prioridad para su aprovacion de tutoria. Esta informacion se maneja en 2 fases por lo que primero se guarda la informacion en un cola de prioridad donde va de 1 a 4 y se ingresa a una pila.
  
    |Rango de nota|Prioridad|
    |-------------|---------|
    |   90 - 100  |   1     |
    |   75 - 89   |   2     |
    |   65 - 74   |   3     |
    |   61 - 64   |   4     |

  * Archivo cursos [.Json]: Se emplea esta opcion la cual permite ingresar un listado de cursos los cuales tienen los datos de nombre y codigo y se van guardando en un ArbolAVL.
  

  * Administrar solicitudes de tutores: aqui se permite al administrador aceptar o rechazar los tutores, aqui se muestran primero los tutores segun su prioridad, los datos se almacenan en una lista doble circular.


  * Reportes: se hace uso de la herramienta de Graphviz para graficar los datos mediante las estructuras que se utilzaron para almacenarlos, como pueden ser listas dobdes, Matrices dispersas, Arboles Avl.
  
  ![Grapviz](https://spin.atomicobject.com/wp-content/uploads/20170927122530/vscode-graphviz-live-preview.gif)



## Estructura de datos utilizadas:

se hizo uso de distintas estructuras las cuales sirvieron para almacenar los datos y asi poder analizarlos y almacenarlos:

* Listas Dobles: Son estructura que se caracterizan por poseer 2 punteros Anterior y Siguiente los caules apuntan al dato respectivo. Ventaja de usar listas es que son estructuras dinamicas lo que quiere decir que no almacenan una cantidad fija de datos, lo cual es importante ya que crece su espacio segun lo que se va consumiento.
![listaD2](https://ccia.ugr.es/~jfv/ed1/tedi/cdrom/icons/lenlaz1.gif)

``` Go
\\Creacion de un nodo y una lista doble en Go

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

* Matrices Dispersas: son estructuras de datos diseñadas para almacenar matrices que contienen una gran cantidad de elementos con valores predeterminados y una pequeña cantidad de elementos distintos de ese valor predeterminado. 

![matriz](https://github.com/CristianMejia2198/EDD_1S_JUNIO_2023/blob/main/Clase6/matriz.jpg?raw=true)

*  Arbol Avl: es un árbol binario de búsqueda  que satisface la condición de estar balanceado. 
  
  ![arbolAvl](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ9PEPUmF14PwYNn-VP9asP7vWyzTOZTlICsA&usqp=CAU)

