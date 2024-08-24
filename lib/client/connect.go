package client

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/pasannissanka/network_go/lib/server"
)

func Connect(conn net.Conn) error {
	fmt.Println("Connecting to server...")

	err := conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	if err != nil {
		fmt.Println("SetReadDeadline failed:", err)
		// do something else, for example create new conn
		return fmt.Errorf("SetReadDeadline failed: %s", err)
	}

	fmt.Fprintf(conn, "Hi UDP Server, How are you doing?")

	// Create a temp buffer
	p := make([]byte, 2048)

	_, err = conn.Read(p)

	fmt.Println("Received message from server: ", string(p))

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
		fmt.Println("Error decoding from server: ", err)
		return fmt.Errorf("error decoding from server: %s", err)
	}

	fmt.Printf("Received message: %v\n", *tmpStruct)

	return connect(*tmpStruct)
}

func connect(message server.Message) error {
	if HasConnection(message.ID) {
		return fmt.Errorf("Connection already exists")
	}

	port, err := strconv.Atoi(message.PORT)
	if err != nil {
		return fmt.Errorf("error converting port to int: %s", err)
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", message.IP, port))

	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		return fmt.Errorf("TCP Resolve failed: %s", err)
	}

	strEcho := "Hello TCP Server!"

	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		return fmt.Errorf("TCP Dial failed: %s", err)
	}

	_, err = tcpConn.Write([]byte(strEcho))
	if err != nil {
		println("Write to server failed:", err.Error())
		return fmt.Errorf("TCP Heartbeat check failed: %s", err)
	}

	println("write to server = ", strEcho)

	reply := make([]byte, 1024)

	_, err = tcpConn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		return fmt.Errorf("TCP Heartbeat check failed: %s", err)
	}

	println("reply from server=", string(reply))

	connection := Connection{
		Conn: tcpConn,
		Id:   message.ID,
		Ip:   message.IP,
		Port: port,
	}

	AddConnection(connection)

	return nil
}
