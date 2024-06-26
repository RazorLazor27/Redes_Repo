package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strings"
)

const Puerto = ":8080"

// Mostrar si el servido ha validado la conexion UDP
var UDPServerStatus bool = true

// Mostrar si el servido ha validado la conexion TCP
var TCPServerStatus bool = false
var Ip string

var codeExit string = ""

// Se explica con el nombre, esta funcion escoge una letra al azar entre las 4 posibles que la tarea nos da
func letra_azar() string {
	letras := []string{"A", "B", "C", "D"}

	indice := rand.Intn(len(letras))

	return letras[indice]

}

// Esta funcion mantiene la conexion UDP abierta mientras sea necesario, ayuda a tener un codigo mas conciso en la funcion main
func cicloCliente(conexion *net.UDPConn) {

	buffer := make([]byte, 1024)
	n, addr, err := conexion.ReadFromUDP(buffer)
	fmt.Print(" --- ", string(buffer[0:n]))

	if err != nil {
		fmt.Println("Error en la conexion, cerrando el servidor", err)
		return
	}

	fmt.Println("El mensaje recibido de: ", addr.String(), ", fue:", string(buffer[:n]))
	Ip = addr.String()
	_, err = conexion.WriteToUDP([]byte("Soy la respuesta del servidor"), addr)
	if err != nil {
		fmt.Println("Error al enviar la respuesta al cliente", err)
		return
	}

	msg := strings.ToUpper(string(buffer[:n]))

	if strings.Contains(msg, "SI") {
		fmt.Println("La conexion UDP fue establecida, procediendo al siguiente protocolo")
		UDPServerStatus = false
		TCPServerStatus = true
	} else {
		fmt.Println("Conexion no aceptada, procediendo a cerrar la conexion UDP")
		UDPServerStatus = false
	}
}

func main() {

	s, err := net.ResolveUDPAddr("udp4", Puerto)

	if err != nil {
		fmt.Println(err)
		return
	}

	conexion, err := net.ListenUDP("udp4", s)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conexion.Close()

	fmt.Println("El servidor UDP esta en funcionamiento desde ahora en el puerto:", Puerto)

	for UDPServerStatus {
		cicloCliente(conexion)
	}

	if TCPServerStatus {

		//fmt.Println("La conexion TCP ocurrira en la direccion ip: ", Ip)
		fmt.Println("Iniciando la conexion TCP...")
		letraJugador := letra_azar()
		letraServer := letra_azar()

		fmt.Println("La letra del jugador al inicio es:", letraJugador)

		l, err := net.Listen("tcp", Puerto)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer l.Close()

		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		testeo := letraJugador
		_, err = c.Write([]byte(testeo + letraServer + "\n"))
		if err != nil {
			fmt.Println("ERROR CTM:", err)
		}

		// tablaArrayS := []string{"A", "B", "C", "D"}

		fmt.Printf("Letra Servidor: %s\n", letraServer)

		var x string

		for {
			networkData, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}

			//tcpmsg viene siendo el mensaje del jugador al servidor

			tcpmsg := strings.TrimSpace(string(networkData))
			tcpmsg = strings.ToUpper(tcpmsg)

			fmt.Println("El mensaje del jugador es:", tcpmsg)
			if tcpmsg == "STOP" {
				fmt.Println("Saliendo del Servidor TCP")
				codeExit = "4"
				x = "F"

			} else if tcpmsg == letraServer {
				//Jugador gana
				fmt.Println("EL JUGADOR GANA")
				codeExit = "1"
				x = "F"

			} else {
				// Aqui va el servidor a simular un disparo
				x = letra_azar()
				fmt.Println("el servidor disparo hacia:", x)
				if x == letraJugador {
					// Servidor gana
					fmt.Println("LA COMPUTADORA HA GANADO")
					codeExit = "3"
				} else {
					// El juego continua
					codeExit = "2"
				}
			}

			fmt.Println("-> ", tcpmsg)
			// t := time.Now()

			// mytime := t.Format(time.RFC3339) + "\n"
			// La funcion write escribe el mensaje hacia el cliente
			// Para simplificar ocuparemos codigos en int para
			// Denotar si el juego sigue o no

			// Listado de numeros:
			/*
				(1) -> El jugador ha encontrado al enemigo (jugador gana)
				(2) -> El jugador ha fallado y el servidor tambien (juego sigue)
				(3) -> El jugador ha fallado pero el servidor no (servidor gana)
				(4) -> El jugador detiene el juego
			*/
			salida := codeExit + x + "\n"

			fmt.Println("El codigo de salida del servidor es:", salida)

			c.Write([]byte(salida))

			if codeExit != "2" {
				return
			}
		}

		// Jugador dice que si
	} else {
		// Jugador dice cualquier cosa pero nosotros lo consideramos como un no
	}

}
