package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

const Puerto = ":8080"

var UDPServerStatus bool = true
var TCPServerStatus bool = false
var Ip string

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

		fmt.Println("La conexion TCP ocurrira en la direccion ip: ", Ip)
		fmt.Println("Iniciando la conexion TCP...")

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

		for {
			networkData, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}

			tcpmsg := strings.TrimSpace(string(networkData))
			tcpmsg = strings.ToUpper(tcpmsg)

			if tcpmsg == "STOP" {
				fmt.Println("Saliendo del Servidor TCP")
				return
			}

			fmt.Println("-> ", tcpmsg)
			t := time.Now()

			mytime := t.Format(time.RFC3339) + "\n"
			c.Write([]byte(mytime))
		}

		// Jugador dice que si
	} else {
		// Jugador dice cualquier cosa pero nosotros lo consideramos como un no
	}

}
