package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var UDPClientStatus bool = true
var TCPClientStatus bool = false

var tableroJ = []string{"A", "B", "C", "D"}
var tableroS = []string{"A", "B", "C", "D"}

func cambiarLetra(letra string, arr []string) {
	for i := range arr {
		if arr[i] == letra {
			arr[i] = "X"
		}
	}
}

func show_tablero(tablero []string) {
	fmt.Println("\n                    Tablero")
	fmt.Printf("               ┌─────┐  ┌─────┐\n")
	fmt.Printf("               │  %s  │  │  %s  │\n", tablero[0], tablero[1])
	fmt.Printf("               └─────┘  └─────┘\n")
	fmt.Printf("               ┌─────┐  ┌─────┐\n")
	fmt.Printf("               │  %s  │  │  %s  │\n", tablero[2], tablero[3])
	fmt.Printf("               └─────┘  └─────┘\n")
}

func main() {
	argumentos := os.Args

	if len(argumentos) == 1 {
		fmt.Println("Favor de proveer un host y su direccion ip")
		return
	}

	conexion := argumentos[1]

	fmt.Println(conexion)

	s, err := net.ResolveUDPAddr("udp4", conexion)

	if err != nil {
		fmt.Println("No resolvio :(", err)
		return
	}

	c, err := net.DialUDP("udp4", nil, s)

	if err != nil {
		fmt.Println("No se logro llegar al servidor", err)
		return
	}

	fmt.Printf("El servidor UDP es %s\n", c.RemoteAddr().String())
	defer c.Close()

	for UDPClientStatus {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(" -> ")
		texto, _ := reader.ReadString('\n')
		data := []byte(texto + "\n")
		_, err := c.Write(data)

		CliMessage := strings.TrimSpace(string(data))

		CliMessage = strings.ToUpper(CliMessage)

		if CliMessage == "SI" {
			fmt.Println("Conexion UDP aprobada, pasando al servidor TCP")
			UDPClientStatus = false
			TCPClientStatus = true
		} else {
			fmt.Println("Conexion UDP rechazada, Cerrando la conexion...")
			UDPClientStatus = false
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		buffer := make([]byte, 1024)

		n, _, err := c.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Respuesta: %s\n", string(buffer[0:n]))

	}

	//Aqui comienza la parte del juego mismo

	if TCPClientStatus {
		//Aqui se probara el servidor TCP y la conexion de mierda

		fmt.Println("Iniciando conexion TCP con el servidor")
		time.Sleep(1 * time.Second)
		fmt.Println("iniciado")

		conexionTCP := argumentos[1]

		conn, err := net.Dial("tcp", conexionTCP)

		if err != nil {
			fmt.Println(err)
			return
		}

		msg, _ := bufio.NewReader(conn).ReadString('\n')

		posJugador := string(msg[0])
		posServer := string(msg[1])

		fmt.Println("El valor del mensaje 1 es:", msg)

		fmt.Println("Capitan!, usted se encuentra en la zona:", posJugador)
		fmt.Println("El enemigo se encuentra en la zona:", posServer)

		for TCPClientStatus {
			// Aqui asumimos que el usuario nos da un caracter bueno
			// De no hacerlo, xd
			show_tablero(tableroJ)

			lector := bufio.NewReader(os.Stdin)
			fmt.Println("Que casilla desea atacar, mi capitan?")
			fmt.Print(" > ")

			letra, _ := lector.ReadString('\n')

			fmt.Fprintf(conn, letra+"\n")

			//mensaje es lo que le manda el servidor al cliente

			mensaje, _ := bufio.NewReader(conn).ReadString('\n')

			fmt.Print("(server) --> " + mensaje)

			letra = strings.TrimSpace(string(letra))
			letra = strings.ToUpper(letra)
			fmt.Println("La letra entregada fue: ", letra)

			decision := string(mensaje[0])
			lugar := string(mensaje[1])

			if strings.Contains(letra, "STOP") {
				fmt.Println("Cerrando la conexion TCP del cliente")
				TCPClientStatus = false
				return
			}
			switch decision {

			case "1": //Usuario gano
				fmt.Println("Le has acertado a BlackBeard!, Sales victorioso del combate")
				fmt.Println("Gracias por jugar :D")
				TCPClientStatus = false
				return
			case "2":
				fmt.Println("Has fallado tu tiro!, pero BlackBeard tambien...")
				cambiarLetra(letra, tableroJ)
				cambiarLetra(lugar, tableroS)

			case "3":
				fmt.Println("Has fallado tu tiro!, pero BlackBeard logra acertar el suyo!")
				fmt.Println("Tu tripulación es baja y tu barco se hunde, PERDISTE")
				TCPClientStatus = false
				return
			}

		}

	}

}
