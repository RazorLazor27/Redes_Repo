package dns

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"net"

	"golang.org/x/net/dns/dnsmessage"
)

const ROOT_SERVERS = "198.41.0.4, 199.9.14.201, 192.33.4.12, 199.7.91.13, 192.203.230.10, 192.5.5.241, 192.112.36.4, 198.97.190.53"

func handlePacket(pc net.PacketConn, addr net.Addr, buf []byte) error {
	return fmt.Errorf("le falta cocinarse")
}

func outgoinDnsQuery(servers []net.IP, question dnsmessage.Question) (*dnsmessage.Parser, *dnsmessage.Header, error) {
	fmt.Printf("Una nueva query del servidor dns para %s, servers: %+v \n", question.Name.String(), servers)

	// Numero al azar para definir el ID del mensaje del header
	verstappen := ^uint16(0)
	random, err := rand.Int(rand.Reader, big.NewInt(int64(verstappen)))
	if err != nil {
		return nil, nil, err
	}

	mensaje := dnsmessage.Message{
		Header: dnsmessage.Header{
			ID:       uint16(random.Int64()),
			Response: false,
			OpCode:   dnsmessage.OpCode(0),
		},
		Questions: []dnsmessage.Question{question},
	}

	buf, err := mensaje.Pack()
	if err != nil {
		return nil, nil, err
	}

	var conn net.Conn

	for _, server := range servers {
		conn, err = net.Dial("udp", server.String()+":53")
		if err != nil {
			break
		}
	}
	if conn == nil {
		return nil, nil, fmt.Errorf("no se logro establecer la conexion a los servidores: %s", err)
	}

	_, err = conn.Write(buf)
	if err != nil {
		return nil, nil, err
	}

	respuesta := make([]byte, 512)

	n, err := bufio.NewReader(conn).Read(respuesta)
	if err != nil {
		return nil, nil, err
	}

	conn.Close()

	var p dnsmessage.Parser

	header, err := p.Start(respuesta[:n])
	if err != nil {
		return nil, nil, fmt.Errorf("el parser fallo su ejecucion: %s", err)
	}

	preguntas, err := p.AllQuestions()
	if err != nil {
		return nil, nil, err
	}

	if len(preguntas) != len(mensaje.Questions) {
		return nil, nil, fmt.Errorf("el packet de respuestas no tiene el mismo tamano que el de preguntas")
	}

	err = p.SkipAllQuestions()
	if err != nil {
		return nil, nil, err
	}

	return &p, &header, nil

}
