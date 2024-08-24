package client

import "net"

type Connection struct {
	Conn net.Conn
	Id   int
	Ip   string
	Port int
}

var (
	Connections []Connection
)
