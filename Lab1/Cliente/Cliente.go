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

		c, err := net.Dial("tcp", conexionTCP)

		if err != nil {
			fmt.Println(err)
			return
		}

		// lectoraux := bufio.NewReader(os.Stdin)
		// letraaux, _ := lectoraux.ReadString('\n')

		fmt.Println("Capitan!, usted se encuentra en la zona:", "A")

		for TCPClientStatus {
			// Aqui asumimos que el usuario nos da un caracter bueno
			// De no hacerlo, xd
			lector := bufio.NewReader(os.Stdin)
			fmt.Println("Que casilla desea atacar, mi capitan?")
			fmt.Print(" > ")

			letra, _ := lector.ReadString('\n')

			fmt.Fprintf(c, letra+"\n")

			//mensaje es lo que le manda el servidor al cliente

			mensaje, _ := bufio.NewReader(c).ReadString('\n')

			fmt.Print("(server) --> " + mensaje)

			letra = strings.TrimSpace(string(letra))
			letra = strings.ToUpper(letra)
			fmt.Println("La letra entregada fue: ", letra)

			if strings.Contains(letra, "STOP") {
				fmt.Println("Cerrando la conexion TCP del cliente")
				TCPClientStatus = false
				return
			}

		}

	}

}
