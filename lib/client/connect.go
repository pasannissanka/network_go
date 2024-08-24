package client

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/pasannissanka/network_go/server"
)

func Connect(conn net.Conn) {
	// Create a temp buffer
	p := make([]byte, 2048)

	conn.Read(p)

	// convert bytes into Buffer (which implements io.Reader/io.Writer)
	tmpBuff := bytes.NewBuffer(p)
	tmpStruct := new(server.Message)

	// creates a decoder object
	gobObjDec := gob.NewDecoder(tmpBuff)
	// decodes buffer and unmarshals it into a Message struct
	gobObjDec.Decode(tmpStruct)

	fmt.Printf("Received response from server: [%s]\n", tmpStruct)

	connect(*tmpStruct)
}

func connect(message server.Message) {
	port, err := strconv.Atoi(message.PORT)

	if err != nil {
		fmt.Println("Error converting port to integer: ", err)
	}

	addr := net.TCPAddr{
		IP:   net.IP(message.IP),
		Port: port,
	}

	strEcho := "Hello TCP Server!"

	tcpConn, err := net.DialTCP("tcp", nil, &addr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = tcpConn.Write([]byte(strEcho))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("write to server = ", strEcho)

	reply := make([]byte, 1024)

	_, err = tcpConn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("reply from server=", string(reply))

	defer tcpConn.Close()
}
