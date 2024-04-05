Nombre: Javiera Osorio Mardones 
Rol: 202173656-5
Paralelo: 201

Nombre: Nicolas Berenguela Pérez 
Rol: 202173619-0
Paralelo: 201

Para la realización de este laboratorio ocupamos el Lenguaje de Programación Go en su version 1.22.2.
El programa fue corrido en 2 maquinas las cuales corrian windows 10 y linux Fedora 39

--> Instrucciones para el uso del programa:

	-> Para la correcta realizacion del programa se van a requerir un par de pasos especiales, lo primero es crear un terminal en la carpeta Lab1 donde
	correra el archivo Server.go usando el siguiente comando "go run Server.go". Luego se necesitara abrir otro terminal nuevo (Donde se realizara el juego)
	Este terminal tiene que estar apuntando a la carpeta Lab1/Cliente, dentro de esa carpeta, en el terminal, se correra el archivo Cliente.go usando el 
	siguiente comando "go run Cliente.go 127.0.0.1:8080". Donde 127.0.0.1 es la dirección ip y :8080 el puerto de conexion.

	-> A pesar de tener 2 terminales abiertos, solo se necesita el terminal que esta corriendo el archivo "Cliente.go", el terminal con el archivo 
	"Server.go" solo muestra por pantalla lo que esta enviando el servidor y algunos detalles que facilitan la victoria dentro del juego

	-> Para jugar el juego el jugador tiene que estar en el terminal corriendo "Cliente.go", lo primero que debera hacer el usuario es escribir por pantalla 
	"si" para poder aprobar la conexion UDP del cliente al servidor, una vez eso haya sido aceptado, comienza el juego. En el el jugador debera escribir
	por pantalla la letra del cuadrante que desee atacar, cuidado si que BlackBeard puede disparar devuelta ;)


--> Consideraciones:

	-> Es de vital importancia seguir el orden puesto arriba de como iniciar nuestro programa, mucho enfasís en que se tiene que correr el Servidor primero
	antes de correr el CLiente
	-> El juego no comprueba si las teclas apretadas coinciden con una letra en el mapa, por lo que si alguien escribe "casa" por ejemplo durante el juego,
	se asumira que escogio una palabra mala y regalara su turno
	-> Nosotros para hacer correr nuestro programa usamos los valores de la direccion ip y puertos dados, importante que se ocupen solo esos valores


