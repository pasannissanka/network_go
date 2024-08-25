package net

import (
	"fmt"
	"log"
	"net"
	"time"
)

type Connection struct {
	Conn net.Conn
	Id   int
	Ip   string
	Port int
	Addr net.Addr
}

func (c *Connection) Connect() error {
	err := c.Conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	if err != nil {
		log.Fatalf("SetReadDeadline failed: %s", err)
		// do something else, for example create new conn
		return fmt.Errorf("SetReadDeadline failed: %s", err)
	}

	err = c.Conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		log.Fatalf("SetWriteDeadline failed: %s", err)
		// do something else, for example create new conn
		return fmt.Errorf("SetWriteDeadline failed: %s", err)
	}

	// Send ping and wait for pong

	go c.Handle()

	return nil
}

func (c *Connection) Handle() {

	defer c.Conn.Close()
}
