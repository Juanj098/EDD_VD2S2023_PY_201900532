# Manaul tecnico
## Descripcion del proyecto
Se desarrollo un software el cual es una plataforma donde se tienen 3 tipos de usuario, Administrador, Estudiante y tutor en el cual se da apoyo en el manejo de libros dedicados a ciertos cursos asi como permite a los tutores compartir publicaciones. se hicieron uso de distintas estructuras de datos las cuales se detallaran en el documento.

Se creo la aplicacion cliente/servidor donde el cliente(frontend) se desarrollo con javaScript haciendo uso de su libreria React y el servidor(Backend) se programo con Go al igual haciendo uso del framework fiber.

* ## FRONTEND 
  El frontend es la parte que proporciona experiencia al usuario, son las pantallas e interfaces que el cliente puede observar y manipular. en este caso seria la interfaz presentada desde el navegador.

  ![login](imgManuales/Captura%20de%20pantalla%202024-01-05%20194522.png)

    En el desarrollo web por parte del Frontend se tienen diversas herramientas las cuales hacen que el desarrollo sea mas agradable y productivo. En estas herramientas encontramos a:
    
    * *Javascript*: Lenguaje de programacion que fue creado con este proposito el cual era darle mayor dinamismo a las paginas web del lado del frontend, a lo largo del tiempo a crecido hasta llegar a ser tambien utilizado del lado del backend.

    * *React*: React es una biblioteca de JavaScript de código abierto para construir interfaces de usuario. Está basada en la componetización de la UI: la interfaz se divide en componentes independientes, que contienen su propio estado. Cuando el estado de un componente cambia, React vuelve a renderizar la interfaz.
    * *Css*: Hojas de Estilo en Cascada (del inglés Cascading Style Sheets) o CSS es el lenguaje de estilos utilizado para describir la presentación de documentos HTML o XML (en-US) (incluyendo varios lenguajes basados en XML como SVG, MathML o XHTML). CSS describe como debe ser renderizado el elemento estructurado en la pantalla, en papel, en el habla o en otros medios.

    ![react](data:image/png;base64,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)

    * ## BACKEND

    El backend se refiere a la parte de un sistema informático o aplicación que gestiona y procesa la lógica, la manipulación de datos y la funcionalidad no visible para el usuario. En un modelo de desarrollo web, el backend es responsable de realizar operaciones como almacenar y recuperar datos de la base de datos, procesar la lógica de negocio, y gestionar la comunicación entre el frontend (la interfaz de usuario) y la base de datos. Utiliza lenguajes de programación y tecnologías específicas para proporcionar servicios y funcionalidades que permiten que una aplicación web funcione de manera eficiente y segura. En resumen, el backend es la "espalda" de una aplicación, encargándose de las tareas que no son directamente visibles para el usuario pero que son fundamentales para el funcionamiento del sistema.
        ![tecnico](imgManuales/Captura%20de%20pantalla%202024-01-05%20195812.png)

    Se utilizan herramientas las cuales son importantes y proporcionan un gran apoyo en su desarrollo :
    * *Go*: También conocido como Golang, es un lenguaje de     programación de código abierto desarrollado por Google. Es un lenguaje de programación compilado, concurrente y de tipado estático diseñado para ser eficiente, fácil de leer y escribir, y adecuado para el desarrollo de software a gran escala.

        En el contexto del backend, Go se utiliza a menudo para construir servidores y aplicaciones que requieren un rendimiento eficiente y concurrente.
  
    *  *Fiber*: se refiere a un framework web para el lenguaje de programación Go (o Golang). Fiber está diseñado para ser rápido, eficiente y fácil de usar.

        ![fiber](data:image/png;base64,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)
  
  * ## ESTRUCTURAS DE DATOS
  *  *Grafo*: Un grafo es una estructura matemática utilizada para representar relaciones entre objetos. Consiste en un conjunto de nodos (también llamados vértices) y un conjunto de aristas (también llamadas arcos), donde cada arista conecta dos nodos y puede tener una dirección (si es dirigido) o no (si es no dirigido). Los grafos son ampliamente utilizados en teoría de grafos y en muchas áreas de la informática, las matemáticas aplicadas y la ciencia en general.
  *  *Arbol B*: Un árbol B es una estructura de datos en forma de árbol que se utiliza comúnmente en bases de datos y sistemas de almacenamiento para organizar y buscar grandes cantidades de datos de manera eficiente. Fue desarrollado por Rudolf Bayer y Edward M. McCreight en 1972.
  *  *Tabla Hash*: Es una estructura de datos que asocia claves con valores. Permite la búsqueda eficiente de un valor asociado con una clave dada. La idea fundamental de una tabla hash es utilizar una función de hash para convertir la clave en un índice dentro de una tabla. Este índice se utiliza luego para acceder directamente a la posición en la que se encuentra el valor asociado con la clave.
  *  *Arbol de Merkle*: También conocido como árbol hash de Merkle, es una estructura de datos utilizada en criptografía y sistemas distribuidos para verificar la integridad de los datos en un conjunto grande de información. Fue propuesto por Ralph Merkle en 1979 y se ha convertido en un componente fundamental de tecnologías como blockchain.