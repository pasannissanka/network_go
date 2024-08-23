package net

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	bin_buf := new(bytes.Buffer)

	data := TCP_PORT
	// create a encoder object
	gobobj := gob.NewEncoder(bin_buf)
	// encode buffer and marshal it into a gob object
	gobobj.Encode(data)

	_, err := conn.WriteToUDP(bin_buf.Bytes(), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func udp() {
	fmt.Printf("Starting UDP server on ip: %s port : %d \n", ServerData.IP, ServerData.UDP_L)

	addr := net.UDPAddr{
		Port: int(ServerData.UDP_L),
		IP:   net.ParseIP(ServerData.IP),
	}

	conn, err := net.ListenUDP("udp4", &addr)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	p := make([]byte, 2048)

	for {
		_, remote_addr, err := conn.ReadFromUDP(p)
		fmt.Printf("Read a message from %v %s \n", remote_addr, p)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}

		sendResponse(conn, remote_addr)
	}
}

func Connect(conn net.Conn) {
	// Create a temp buffer
	p := make([]byte, 2048)

	conn.Read(p)

	// convert bytes into Buffer (which implements io.Reader/io.Writer)
	tmpBuff := bytes.NewBuffer(p)
	tmpStruct := new(Message)

	// creates a decoder object
	gobObjDec := gob.NewDecoder(tmpBuff)
	// decodes buffer and unmarshals it into a Message struct
	gobObjDec.Decode(tmpStruct)

	fmt.Printf("Received response from server: [%s]\n", tmpStruct)

	defer conn.Close()
}
