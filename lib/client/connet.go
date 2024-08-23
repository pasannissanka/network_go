package client

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"

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

	defer conn.Close()
}
