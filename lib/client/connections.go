package client

import (
	"fmt"
	"net"
)

type Connection struct {
	Conn net.Conn
	Id   int
	Ip   string
	Port int
}

var (
	connections []Connection
)

func AddConnection(conn Connection) {
	fmt.Printf("Adding connection: %v\n", conn)
	connections = append(connections, conn)
}

func RemoveConnection(id int) {
	for i, conn := range connections {
		if conn.Id == id {
			fmt.Printf("Removing connection: %v\n", conn)
			connections = append(connections[:i], connections[i+1:]...)
		}
	}
}

func GetConnection(id int) (connection Connection, err error) {
	for _, conn := range connections {
		if conn.Id == id {
			fmt.Printf("Getting connection: %v\n", conn)
			return conn, nil
		}
	}
	return Connection{}, fmt.Errorf("Connection not found")
}

func HasConnection(id int) bool {
	for _, conn := range connections {
		if conn.Id == id {
			return true
		}
	}
	return false
}

func GetConnections() []Connection {
	return connections
}
