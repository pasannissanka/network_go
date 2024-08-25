package client

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	server_net "github.com/pasannissanka/network_go/lib/net"
	"github.com/pasannissanka/network_go/lib/server"
)

var TcpConnections server_net.Connections

func Connect(conn net.Conn) error {
	err := conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	if err != nil {
		log.Fatalf("SetReadDeadline failed: %s", err)
		// do something else, for example create new conn
		return fmt.Errorf("SetReadDeadline failed: %s", err)
	}

	// Send a message to server
	fmt.Fprintf(conn, "Hi UDP Server, How are you doing?")

	// Create a temp buffer
	p := make([]byte, 2048)

	_, err = conn.Read(p)

	if err != nil {
		return fmt.Errorf("error reading from server: %s", err)
	}

	// convert bytes into Buffer (which implements io.Reader/io.Writer)
	tmpBuff := bytes.NewBuffer(p)
	tmpStruct := new(server.Message)

	// creates a decoder object
	gobObjDec := gob.NewDecoder(tmpBuff)
	// decodes buffer and unmarshals it into a Message struct
	err = gobObjDec.Decode(tmpStruct)

	if err != nil {
		log.Println("Error decoding from server: ", err)
		return fmt.Errorf("error decoding from server: %s", err)
	}

	log.Printf("Received message: %v\n", *tmpStruct)

	return connect(*tmpStruct)
}

func connect(message server.Message) error {
	log.Printf("Received handshake from Node %+v", message)

	if ID == message.ID {
		return fmt.Errorf("connection to self")
	}

	if TcpConnections.HasConnection(message.ID) {
		return fmt.Errorf("connection already exists")
	}

	port, err := strconv.Atoi(message.PORT)
	if err != nil {
		return fmt.Errorf("error converting port to int: %s", err)
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", message.IP, port))
	if err != nil {
		return fmt.Errorf("TCP Resolve failed: %s", err)
	}

	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return fmt.Errorf("TCP Dial failed: %s", err)
	}

	strEcho := "Hello TCP Server!"
	log.Printf("write to server = %s", strEcho)
	_, err = tcpConn.Write([]byte(strEcho))
	if err != nil {
		return fmt.Errorf("TCP Heartbeat check failed: %s", err)
	}

	reply := make([]byte, 1024)
	_, err = tcpConn.Read(reply)
	if err != nil {
		return fmt.Errorf("TCP Heartbeat check failed: %s", err)
	}

	log.Printf("reply from server= %s", string(reply))

	connection := server_net.Connection{
		Conn: tcpConn,
		Id:   message.ID,
		Ip:   message.IP,
		Port: port,
		Addr: tcpAddr,
	}

	TcpConnections.AddConnection(connection)

	return nil
}
